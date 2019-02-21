package bootkube

import (
	"os"
	"path/filepath"

	"github.com/openshift-metalkube/kni-installer/pkg/asset"
	"github.com/openshift-metalkube/kni-installer/pkg/asset/templates/content"
)

const (
	kubeSystemSecretEtcdClientCADeprecatedFileName = "kube-system-secret-etcd-client-ca-deprecated.yaml.template"
)

var _ asset.WritableAsset = (*KubeSystemSecretEtcdClientCADeprecated)(nil)

// KubeSystemSecretEtcdClientCADeprecated is the constant to represent contents of kube-system-secret-etcd-client-ca-deprecated.yaml.template file.
type KubeSystemSecretEtcdClientCADeprecated struct {
	FileList []*asset.File
}

// Dependencies returns all of the dependencies directly needed by the asset
func (t *KubeSystemSecretEtcdClientCADeprecated) Dependencies() []asset.Asset {
	return []asset.Asset{}
}

// Name returns the human-friendly name of the asset.
func (t *KubeSystemSecretEtcdClientCADeprecated) Name() string {
	return "KubeSystemSecretEtcdClientCADeprecated"
}

// Generate generates the actual files by this asset
func (t *KubeSystemSecretEtcdClientCADeprecated) Generate(parents asset.Parents) error {
	fileName := kubeSystemSecretEtcdClientCADeprecatedFileName
	data, err := content.GetBootkubeTemplate(fileName)
	if err != nil {
		return err
	}
	t.FileList = []*asset.File{
		{
			Filename: filepath.Join(content.TemplateDir, fileName),
			Data:     []byte(data),
		},
	}
	return nil
}

// Files returns the files generated by the asset.
func (t *KubeSystemSecretEtcdClientCADeprecated) Files() []*asset.File {
	return t.FileList
}

// Load returns the asset from disk.
func (t *KubeSystemSecretEtcdClientCADeprecated) Load(f asset.FileFetcher) (bool, error) {
	file, err := f.FetchByName(filepath.Join(content.TemplateDir, kubeSystemSecretEtcdClientCADeprecatedFileName))
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	t.FileList = []*asset.File{file}
	return true, nil
}
