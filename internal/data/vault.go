// Copyright 2023 Nautes Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package data

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"net/http"

	kratoshttp "github.com/go-kratos/kratos/v2/transport/http"
	vault "github.com/hashicorp/vault/api"
	auth "github.com/hashicorp/vault/api/auth/kubernetes"
	commonv1 "github.com/nautes-labs/api-server/api/common/v1"
	"github.com/nautes-labs/api-server/internal/biz"
	kubernetes "github.com/nautes-labs/api-server/pkg/kubernetes"
	vaultproxyv1 "github.com/nautes-labs/api-server/pkg/vaultproxy/v1"
	nautesconfigs "github.com/nautes-labs/pkg/pkg/nautesconfigs"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
)

const (
	_USERNAME              = "default"
	_PERMISSION            = "readonly"
	_ACCESS_TYPE           = "deploykey"
	_DefaultServiceAccount = "api-server-manager"
	_CAPATH                = "/usr/local/share/ca-certificates"
	_VAULTTOKENKEY         = "token"
)

type vaultRepo struct {
	secret vaultproxyv1.SecretHTTPClient
	auth   vaultproxyv1.AuthGrantHTTPClient
	config *nautesconfigs.Config
}

func NewVaultClient(config *nautesconfigs.Config) (biz.Secretrepo, error) {
	http, err := NewHttpClientForVault(config.Secret.Vault.ProxyAddr)
	if err != nil {
		return nil, err
	}

	secret := vaultproxyv1.NewSecretHTTPClient(http)
	auth := vaultproxyv1.NewAuthGrantHTTPClient(http)

	return &vaultRepo{secret: secret, auth: auth, config: config}, nil
}

func NewHttpClient(ca string) (*http.Client, error) {
	if ca == "" {
		return nil, fmt.Errorf("failed to get vault cert")
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM([]byte(ca))
	return &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs: caCertPool,
			},
		},
	}, nil
}

func NewKubernetesAuth(mountPath, token string, roles map[string]string) (*auth.KubernetesAuth, error) {
	if mountPath == "" {
		return nil, fmt.Errorf("failed to get vault mount path")
	}

	role, ok := roles["Api"]
	if !ok {
		return nil, fmt.Errorf("failed to get argo-operator role in nautes config")
	}

	k8sAuth, err := auth.NewKubernetesAuth(
		role,
		auth.WithServiceAccountToken(token),
		auth.WithMountPath(mountPath),
	)

	if err != nil {
		return nil, fmt.Errorf("unable to initialize Kubernetes auth method: %w", err)
	}

	return k8sAuth, nil
}

func GetToken(namespace string) (string, error) {
	sa := &corev1.ServiceAccount{}
	saNamespaceName := types.NamespacedName{
		Namespace: namespace,
		Name:      _DefaultServiceAccount,
	}

	client, err := kubernetes.NewClient()
	if err != nil {
		return "", err
	}

	err = client.Get(context.Background(), saNamespaceName, sa)
	if err != nil {
		return "", err
	}

	secretName := sa.Secrets[0].Name
	secret := &corev1.Secret{}
	secretNamespaceName := types.NamespacedName{
		Namespace: namespace,
		Name:      secretName,
	}

	err = client.Get(context.Background(), secretNamespaceName, secret)
	if err != nil {
		return "", err
	}
	return string(secret.Data[_VAULTTOKENKEY]), nil
}

func (v *vaultRepo) NewVaultClient(ctx context.Context) (*vault.Client, error) {
	httpClient, err := NewHttpClient(v.config.Secret.Vault.CABundle)
	if err != nil {
		return nil, err
	}

	token, err := GetToken(v.config.Nautes.Namespace)
	if err != nil {
		return nil, err
	}

	kubernetesAuth, err := NewKubernetesAuth(v.config.Secret.Vault.MountPath, token, v.config.Secret.OperatorName)
	if err != nil {
		return nil, err
	}

	vaultConfig := vault.DefaultConfig()
	vaultConfig.Address = v.config.Secret.Vault.Addr
	vaultConfig.HttpClient = httpClient

	client, err := vault.NewClient(vaultConfig)
	if err != nil {
		return nil, err
	}

	authInfo, err := client.Auth().Login(context.Background(), kubernetesAuth)
	if err != nil {
		return nil, fmt.Errorf("unable to log in with Kubernetes auth: %w", err)
	}

	if authInfo == nil {
		return nil, fmt.Errorf("no auth info was returned after login")
	}

	return client, nil
}

func (v *vaultRepo) Logout(client *vault.Client) error {
	err := client.Auth().Token().RevokeSelf("")
	if err != nil {
		return err
	}
	return nil
}

func (v *vaultRepo) GetDeployKey(ctx context.Context, secretOptions *biz.SecretOptions) (*biz.DeployKeySecretData, error) {
	client, err := v.NewVaultClient(ctx)
	if err != nil {
		return nil, err
	}

	defer func() error {
		err = client.Auth().Token().RevokeSelf("")
		if err != nil {
			return err
		}
		return nil
	}()

	secret, err := client.KVv2(secretOptions.SecretEngine).Get(context.Background(), secretOptions.SecretPath)
	if err != nil {
		err = errors.Unwrap(err)
		if err == vault.ErrSecretNotFound {
			return nil, commonv1.ErrorSecretNotFound("unable to read secret, err: %s", err)
		}

		return nil, fmt.Errorf("unable to read secret: %w", err)
	}

	val, ok := secret.Data[biz.Fingerprint]
	if !ok {
		return nil, commonv1.ErrorSecretNotFound("the fingerprint of the deploy key was not found, secret path: %s", secretOptions.SecretPath)
	}
	if val.(string) == "" {
		return nil, commonv1.ErrorSecretNotFound("the fingerprint of the deploy key was not found, secret path: %s", secretOptions.SecretPath)
	}
	fingerprint := val.(string)

	val, ok = secret.Data[biz.DeployKeyID]
	if !ok {
		return nil, commonv1.ErrorSecretNotFound("the id of the deploy key was not found, secret path: %s", secretOptions.SecretPath)
	}
	if val.(string) == "" {
		return nil, commonv1.ErrorSecretNotFound("the id of the deploy key was not found, secret path: %s", secretOptions.SecretPath)
	}
	id, err := strconv.Atoi(val.(string))
	if err != nil {
		return nil, err
	}

	return &biz.DeployKeySecretData{
		ID:          id,
		Fingerprint: fingerprint,
	}, nil
}

func (v *vaultRepo) GetSecret(ctx context.Context, secretOptions *biz.SecretOptions) (string, error) {
	client, err := v.NewVaultClient(ctx)
	if err != nil {
		return "", err
	}

	defer func() error {
		err = client.Auth().Token().RevokeSelf("")
		if err != nil {
			return err
		}
		return nil
	}()

	secret, err := client.KVv2(secretOptions.SecretEngine).Get(context.Background(), secretOptions.SecretPath)
	if err != nil {
		err = errors.Unwrap(err)
		if err == vault.ErrSecretNotFound {
			return "", commonv1.ErrorSecretNotFound("unable to read secret, err: %s", err)
		}

		return "", fmt.Errorf("unable to read secret: %w", err)
	}

	val, ok := secret.Data[secretOptions.SecretKey]
	if !ok {
		return "", fmt.Errorf("%s secret is not found", secretOptions.SecretKey)
	}
	if val.(string) == "" {
		return "", fmt.Errorf("%s secret data is empty", secretOptions.SecretKey)
	}

	return val.(string), nil
}

func (v *vaultRepo) SaveDeployKey(ctx context.Context, id, key, user, permission string, extendKVs map[string]string) error {
	opt := &vaultproxyv1.GitRequest{
		Meta: &vaultproxyv1.GitMeta{
			ProviderType: string(v.config.Git.GitType),
			Id:           id,
			Username:     user,
			Permission:   permission,
		},
		Kvs: &vaultproxyv1.GitKVs{
			DeployKey:   string(key),
			Additionals: extendKVs,
		},
	}
	_, err := v.secret.CreateGit(context.Background(), opt)
	if err != nil {
		return err
	}

	return nil
}

func (v *vaultRepo) GetProjectAccessToken(ctx context.Context, secretOptions *biz.SecretOptions) (*biz.AccessTokenSecretData, error) {
	client, err := v.NewVaultClient(ctx)
	if err != nil {
		return nil, err
	}

	defer func() error {
		err = client.Auth().Token().RevokeSelf("")
		if err != nil {
			return err
		}
		return nil
	}()

	secret, err := client.KVv2(secretOptions.SecretEngine).Get(context.Background(), secretOptions.SecretPath)
	if err != nil {
		err = errors.Unwrap(err)
		if err == vault.ErrSecretNotFound {
			return nil, commonv1.ErrorSecretNotFound("unable to read secret, err: %s", err)
		}

		return nil, fmt.Errorf("unable to read secret: %w", err)
	}

	val, ok := secret.Data[biz.AccessTokenID]
	if !ok {
		return nil, commonv1.ErrorSecretNotFound("the id of the project access token was not found, secret path: %s", secretOptions.SecretPath)
	}
	if val.(string) == "" {
		return nil, commonv1.ErrorSecretNotFound("the id of the project access token was not found, secret path: %s", secretOptions.SecretPath)
	}
	id, err := strconv.Atoi(val.(string))
	if err != nil {
		return nil, err
	}

	token, ok := secret.Data[biz.SecretsAccessToken]
	if !ok {
		return nil, commonv1.ErrorSecretNotFound("the vault of project access token was not found, secret path: %s", secretOptions.SecretPath)
	}
	if token.(string) == "" {
		return nil, commonv1.ErrorSecretNotFound("the vault of project access token was not found, secret path: %s", secretOptions.SecretPath)
	}

	return &biz.AccessTokenSecretData{
		ID: id,
	}, nil
}

func (v *vaultRepo) SaveProjectAccessToken(ctx context.Context, id, key, user, permission string, extendKVs map[string]string) error {
	opt := &vaultproxyv1.GitRequest{
		Meta: &vaultproxyv1.GitMeta{
			ProviderType: string(v.config.Git.GitType),
			Id:           id,
			Username:     user,
			Permission:   permission,
		},
		Kvs: &vaultproxyv1.GitKVs{
			AccessToken: string(key),
			Additionals: extendKVs,
		},
	}
	_, err := v.secret.CreateGit(context.Background(), opt)
	if err != nil {
		return err
	}

	return nil
}

func (v *vaultRepo) SaveClusterConfig(ctx context.Context, id, config string) error {
	var clustertype = "kubernetes"
	var permission = "admin"
	opt := &vaultproxyv1.ClusterRequest{
		Meta: &vaultproxyv1.ClusterMeta{
			Id:         id,
			Type:       clustertype,
			Username:   _USERNAME,
			Permission: permission,
		},
		Account: &vaultproxyv1.ClusterAccount{
			Kubeconfig: config,
		},
	}
	_, err := v.secret.CreateCluster(context.Background(), opt)
	if err != nil {
		return err
	}

	return nil
}

func (v *vaultRepo) DeleteSecret(ctx context.Context, id int, user, permission string) error {
	repoID := fmt.Sprintf("%s%d", biz.RepoPrefix, id)
	opt := &vaultproxyv1.GitRequest{
		Meta: &vaultproxyv1.GitMeta{
			ProviderType: string(v.config.Git.GitType),
			Id:           repoID,
			Username:     user,
			Permission:   permission,
		},
		// TODO:
		// This is a bug, Subsequent deletion required.
		Kvs: &vaultproxyv1.GitKVs{
			DeployKey: "deploy_key",
		},
	}
	_, err := v.secret.DeleteGit(context.TODO(), opt)
	if err != nil {
		return err
	}
	return nil
}

func (v *vaultRepo) AuthorizationSecret(ctx context.Context, id int, destUser, gitType, mountPath string) error {
	if id == 0 || destUser == "" {
		return fmt.Errorf("authorization failed. please check the parameters")
	}

	destUser, ok := v.config.Secret.OperatorName[destUser]
	if !ok {
		return fmt.Errorf("dest user is not found")
	}

	repoID := fmt.Sprintf("%s%d", biz.RepoPrefix, id)
	opt := &vaultproxyv1.AuthroleGitPolicyRequest{
		ClusterName: mountPath,
		DestUser:    destUser,
		Secret: &vaultproxyv1.GitMeta{
			ProviderType: gitType,
			Id:           repoID,
			Username:     _USERNAME,
			Permission:   _PERMISSION,
		},
	}
	_, err := v.auth.GrantAuthroleGitPolicy(context.Background(), opt)
	if err != nil {
		return err
	}
	return nil
}

func NewHttpClientForVault(serverAddress string) (*kratoshttp.Client, error) {
	content, err := url.Parse(serverAddress)
	if err != nil {
		return nil, err
	}

	host := content.Host
	splits := strings.Split(host, ":")
	if len(splits) == 2 && len(splits[1]) != 0 {
		host = splits[0]
	}

	caCertFilePath := fmt.Sprintf("%v/ca.crt", _CAPATH)
	apiServerCrtFilePath := fmt.Sprintf("%v/apiserver.crt", _CAPATH)
	apiServerCrtKeyPath := fmt.Sprintf("%v/apiserver.key", _CAPATH)

	caCert, err := os.ReadFile(caCertFilePath)
	if err != nil {
		return nil, err
	}
	cert, err := tls.LoadX509KeyPair(apiServerCrtFilePath, apiServerCrtKeyPath)
	if err != nil {
		return nil, err
	}

	cp := x509.NewCertPool()
	if !cp.AppendCertsFromPEM(caCert) {
		return nil, err
	}

	tlsConf := &tls.Config{
		ServerName:   host,
		RootCAs:      cp,
		Certificates: []tls.Certificate{cert},
	}

	conn, err := kratoshttp.NewClient(context.Background(), kratoshttp.WithEndpoint(serverAddress), kratoshttp.WithTLSConfig(tlsConf), kratoshttp.WithTimeout(3*time.Second))
	if err != nil {
		return nil, err
	}

	return conn, nil
}
