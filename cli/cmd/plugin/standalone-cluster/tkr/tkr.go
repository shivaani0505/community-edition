// Copyright 2021 VMware Tanzu Community Edition contributors. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Package tkr contains functions for working with Tanzu Kubernetes Release information.
package tkr

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

const (
	tceRepoURL = "projects.registry.vmware.com/tce/main:0.9.1"
)

//nolint:golint
type TKRBom struct {
	APIVersion string `yaml:"apiVersion"`
	Release    struct {
		Version string `yaml:"version"`
	} `yaml:"release"`
	Components struct {
		AkoOperator []struct {
			Version string `yaml:"version"`
			Images  struct {
				AkoOperatorImage struct {
					ImagePath   string `yaml:"imagePath"`
					Tag         string `yaml:"tag"`
					Reposoitory string `yaml:"repository"`
				} `yaml:"akoOperatorImage"`
			} `yaml:"images"`
		} `yaml:"ako-operator"`
		Antrea []struct {
			Version string `yaml:"version"`
			Images  struct {
				AntreaImage struct {
					ImagePath   string `yaml:"imagePath"`
					Tag         string `yaml:"tag"`
					Reposoitory string `yaml:"repository"`
				} `yaml:"antreaImage"`
			} `yaml:"images"`
		} `yaml:"antrea"`
		CalicoAll []struct {
			Version string `yaml:"version"`
			Images  struct {
				CalicoCniImage struct {
					ImagePath   string `yaml:"imagePath"`
					Tag         string `yaml:"tag"`
					Reposoitory string `yaml:"repository"`
				} `yaml:"calicoCniImage"`
				CalicoKubecontrollerImage struct {
					ImagePath   string `yaml:"imagePath"`
					Tag         string `yaml:"tag"`
					Reposoitory string `yaml:"repository"`
				} `yaml:"calicoKubecontrollerImage"`
				CalicoNodeImage struct {
					ImagePath   string `yaml:"imagePath"`
					Tag         string `yaml:"tag"`
					Reposoitory string `yaml:"repository"`
				} `yaml:"calicoNodeImage"`
				CalicoPodDaemonImage struct {
					ImagePath   string `yaml:"imagePath"`
					Tag         string `yaml:"tag"`
					Reposoitory string `yaml:"repository"`
				} `yaml:"calicoPodDaemonImage"`
			} `yaml:"images"`
		} `yaml:"calico_all"`
		CloudProviderVsphere []struct {
			Version string `yaml:"version"`
			Images  struct {
				CcmControllerImage struct {
					ImagePath   string `yaml:"imagePath"`
					Tag         string `yaml:"tag"`
					Reposoitory string `yaml:"repository"`
				} `yaml:"ccmControllerImage"`
			} `yaml:"images"`
		} `yaml:"cloud_provider_vsphere"`
		CniPlugins []struct {
			Version string `yaml:"version"`
		} `yaml:"cni_plugins"`
		Containerd []struct {
			Version string `yaml:"version"`
		} `yaml:"containerd"`
		Coredns []struct {
			Version string `yaml:"version"`
			Images  struct {
				Coredns struct {
					ImagePath   string `yaml:"imagePath"`
					Tag         string `yaml:"tag"`
					Reposoitory string `yaml:"repository"`
				} `yaml:"coredns"`
			} `yaml:"images"`
		} `yaml:"coredns"`
		CriTools []struct {
			Version string `yaml:"version"`
		} `yaml:"cri_tools"`
		CsiAttacher []struct {
			Version string `yaml:"version"`
			Images  struct {
				CsiAttacherImage struct {
					ImagePath   string `yaml:"imagePath"`
					Tag         string `yaml:"tag"`
					Reposoitory string `yaml:"repository"`
				} `yaml:"csiAttacherImage"`
			} `yaml:"images"`
		} `yaml:"csi_attacher"`
		CsiLivenessprobe []struct {
			Version string `yaml:"version"`
			Images  struct {
				CsiLivenessProbeImage struct {
					ImagePath   string `yaml:"imagePath"`
					Tag         string `yaml:"tag"`
					Reposoitory string `yaml:"repository"`
				} `yaml:"csiLivenessProbeImage"`
			} `yaml:"images"`
		} `yaml:"csi_livenessprobe"`
		CsiNodeDriverRegistrar []struct {
			Version string `yaml:"version"`
			Images  struct {
				CsiNodeDriverRegistrarImage struct {
					ImagePath   string `yaml:"imagePath"`
					Tag         string `yaml:"tag"`
					Reposoitory string `yaml:"repository"`
				} `yaml:"csiNodeDriverRegistrarImage"`
			} `yaml:"images"`
		} `yaml:"csi_node_driver_registrar"`
		CsiProvisioner []struct {
			Version string `yaml:"version"`
			Images  struct {
				CsiProvisonerImage struct {
					ImagePath   string `yaml:"imagePath"`
					Tag         string `yaml:"tag"`
					Reposoitory string `yaml:"repository"`
				} `yaml:"csiProvisonerImage"`
			} `yaml:"images"`
		} `yaml:"csi_provisioner"`
		Dex []struct {
			Version string `yaml:"version"`
			Images  struct {
				DexImage struct {
					ImagePath   string `yaml:"imagePath"`
					Tag         string `yaml:"tag"`
					Reposoitory string `yaml:"repository"`
				} `yaml:"dexImage"`
			} `yaml:"images"`
		} `yaml:"dex"`
		Etcd []struct {
			Version string `yaml:"version"`
			Images  struct {
				Etcd struct {
					ImagePath   string `yaml:"imagePath"`
					Tag         string `yaml:"tag"`
					Reposoitory string `yaml:"repository"`
				} `yaml:"etcd"`
			} `yaml:"images"`
		} `yaml:"etcd"`
		KappController []struct {
			Version string `yaml:"version"`
			Images  struct {
				KappControllerImage struct {
					ImagePath   string `yaml:"imagePath"`
					Tag         string `yaml:"tag"`
					Reposoitory string `yaml:"repository"`
				} `yaml:"kappControllerImage"`
			} `yaml:"images"`
		} `yaml:"kapp-controller"`
		Kubernetes []struct {
			Version string `yaml:"version"`
			Images  struct {
				KubeAPIServer struct {
					ImagePath   string `yaml:"imagePath"`
					Tag         string `yaml:"tag"`
					Reposoitory string `yaml:"repository"`
				} `yaml:"kubeAPIServer"`
				KubeControllerManager struct {
					ImagePath   string `yaml:"imagePath"`
					Tag         string `yaml:"tag"`
					Reposoitory string `yaml:"repository"`
				} `yaml:"kubeControllerManager"`
				KubeE2E struct {
					ImagePath   string `yaml:"imagePath"`
					Tag         string `yaml:"tag"`
					Reposoitory string `yaml:"repository"`
				} `yaml:"kubeE2e"`
				KubeProxy struct {
					ImagePath   string `yaml:"imagePath"`
					Tag         string `yaml:"tag"`
					Reposoitory string `yaml:"repository"`
				} `yaml:"kubeProxy"`
				KubeScheduler struct {
					ImagePath   string `yaml:"imagePath"`
					Tag         string `yaml:"tag"`
					Reposoitory string `yaml:"repository"`
				} `yaml:"kubeScheduler"`
				Pause struct {
					ImagePath   string `yaml:"imagePath"`
					Tag         string `yaml:"tag"`
					Reposoitory string `yaml:"repository"`
				} `yaml:"pause"`
				PauseWindows1809 struct {
					ImagePath   string `yaml:"imagePath"`
					Tag         string `yaml:"tag"`
					Reposoitory string `yaml:"repository"`
				} `yaml:"pause_windows_1809"`
				PauseWindows1903 struct {
					ImagePath   string `yaml:"imagePath"`
					Tag         string `yaml:"tag"`
					Reposoitory string `yaml:"repository"`
				} `yaml:"pause_windows_1903"`
				PauseWindows1909 struct {
					ImagePath   string `yaml:"imagePath"`
					Tag         string `yaml:"tag"`
					Reposoitory string `yaml:"repository"`
				} `yaml:"pause_windows_1909"`
				PauseWindows2004 struct {
					ImagePath   string `yaml:"imagePath"`
					Tag         string `yaml:"tag"`
					Reposoitory string `yaml:"repository"`
				} `yaml:"pause_windows_2004"`
			} `yaml:"images"`
		} `yaml:"kubernetes"`
		KubernetesCsiExternalResizer []struct {
			Version string `yaml:"version"`
			Images  struct {
				CsiExternalResizer struct {
					ImagePath   string `yaml:"imagePath"`
					Tag         string `yaml:"tag"`
					Reposoitory string `yaml:"repository"`
				} `yaml:"csiExternalResizer"`
			} `yaml:"images"`
		} `yaml:"kubernetes-csi_external-resizer"`
		KubernetesSigsKind []struct {
			Version string `yaml:"version"`
			Images  struct {
				KindNodeImage struct {
					ImagePath   string `yaml:"imagePath"`
					Tag         string `yaml:"tag"`
					Reposoitory string `yaml:"repository"`
				} `yaml:"kindNodeImage"`
			} `yaml:"images"`
		} `yaml:"kubernetes-sigs_kind"`
		LoadBalancerAndIngressService []struct {
			Version string `yaml:"version"`
			Images  struct {
				LoadBalancerAndIngressServiceImage struct {
					ImagePath   string `yaml:"imagePath"`
					Tag         string `yaml:"tag"`
					Reposoitory string `yaml:"repository"`
				} `yaml:"loadBalancerAndIngressServiceImage"`
			} `yaml:"images"`
		} `yaml:"load-balancer-and-ingress-service"`
		MetricsServer []struct {
			Version string `yaml:"version"`
			Images  struct {
				MetricsServerImage struct {
					ImagePath   string `yaml:"imagePath"`
					Tag         string `yaml:"tag"`
					Reposoitory string `yaml:"repository"`
				} `yaml:"metricsServerImage"`
			} `yaml:"images"`
		} `yaml:"metrics-server"`
		Pinniped []struct {
			Version string `yaml:"version"`
			Images  struct {
				PinnipedImage struct {
					ImagePath   string `yaml:"imagePath"`
					Tag         string `yaml:"tag"`
					Reposoitory string `yaml:"repository"`
				} `yaml:"pinnipedImage"`
			} `yaml:"images"`
		} `yaml:"pinniped"`
		TanzuFrameworkAddons []struct {
			Version string `yaml:"version"`
			Images  struct {
				TanzuAddonsManagerImage struct {
					ImagePath   string `yaml:"imagePath"`
					Tag         string `yaml:"tag"`
					Reposoitory string `yaml:"repository"`
				} `yaml:"tanzuAddonsManagerImage"`
				TkgPinnipedPostDeployImage struct {
					ImagePath   string `yaml:"imagePath"`
					Tag         string `yaml:"tag"`
					Reposoitory string `yaml:"repository"`
				} `yaml:"tkgPinnipedPostDeployImage"`
			} `yaml:"images"`
		} `yaml:"tanzu-framework-addons"`
		TkgCorePackages []struct {
			Version string `yaml:"version"`
			Images  struct {
				AddonsManagerTanzuVmwareCom struct {
					ImagePath   string `yaml:"imagePath"`
					Tag         string `yaml:"tag"`
					Reposoitory string `yaml:"repository"`
				} `yaml:"addons-manager.tanzu.vmware.com"`
				AkoOperatorTanzuVmwareCom struct {
					ImagePath   string `yaml:"imagePath"`
					Tag         string `yaml:"tag"`
					Reposoitory string `yaml:"repository"`
				} `yaml:"ako-operator.tanzu.vmware.com"`
				AntreaTanzuVmwareCom struct {
					ImagePath   string `yaml:"imagePath"`
					Tag         string `yaml:"tag"`
					Reposoitory string `yaml:"repository"`
				} `yaml:"antrea.tanzu.vmware.com"`
				CalicoTanzuVmwareCom struct {
					ImagePath   string `yaml:"imagePath"`
					Tag         string `yaml:"tag"`
					Reposoitory string `yaml:"repository"`
				} `yaml:"calico.tanzu.vmware.com"`
				KappControllerTanzuVmwareCom struct {
					ImagePath   string `yaml:"imagePath"`
					Tag         string `yaml:"tag"`
					Reposoitory string `yaml:"repository"`
				} `yaml:"kapp-controller.tanzu.vmware.com"`
				LoadBalancerAndIngressServiceTanzuVmwareCom struct {
					ImagePath   string `yaml:"imagePath"`
					Tag         string `yaml:"tag"`
					Reposoitory string `yaml:"repository"`
				} `yaml:"load-balancer-and-ingress-service.tanzu.vmware.com"`
				MetricsServerTanzuVmwareCom struct {
					ImagePath   string `yaml:"imagePath"`
					Tag         string `yaml:"tag"`
					Reposoitory string `yaml:"repository"`
				} `yaml:"metrics-server.tanzu.vmware.com"`
				PinnipedTanzuVmwareCom struct {
					ImagePath   string `yaml:"imagePath"`
					Tag         string `yaml:"tag"`
					Reposoitory string `yaml:"repository"`
				} `yaml:"pinniped.tanzu.vmware.com"`
				TanzuCorePackageRepositoryImage struct {
					ImagePath   string `yaml:"imagePath"`
					Tag         string `yaml:"tag"`
					Reposoitory string `yaml:"repository"`
				} `yaml:"tanzuCorePackageRepositoryImage"`
				VsphereCpiTanzuVmwareCom struct {
					ImagePath   string `yaml:"imagePath"`
					Tag         string `yaml:"tag"`
					Reposoitory string `yaml:"repository"`
				} `yaml:"vsphere-cpi.tanzu.vmware.com"`
				VsphereCsiTanzuVmwareCom struct {
					ImagePath   string `yaml:"imagePath"`
					Tag         string `yaml:"tag"`
					Reposoitory string `yaml:"repository"`
				} `yaml:"vsphere-csi.tanzu.vmware.com"`
			} `yaml:"images"`
		} `yaml:"tkg-core-packages"`
		VsphereCsiDriver []struct {
			Version string `yaml:"version"`
			Images  struct {
				CsiControllerImage struct {
					ImagePath string `yaml:"imagePath"`
					Tag       string `yaml:"tag"`
				} `yaml:"csiControllerImage"`
				CsiMetaDataSyncerImage struct {
					ImagePath string `yaml:"imagePath"`
					Tag       string `yaml:"tag"`
				} `yaml:"csiMetaDataSyncerImage"`
			} `yaml:"images"`
		} `yaml:"vsphere_csi_driver"`
	} `yaml:"components"`
	KubeadmConfigSpec struct {
		APIVersion        string `yaml:"apiVersion"`
		Kind              string `yaml:"kind"`
		ImageRepository   string `yaml:"imageRepository"`
		KubernetesVersion string `yaml:"kubernetesVersion"`
		Etcd              struct {
			Local struct {
				DataDir         string `yaml:"dataDir"`
				ImageRepository string `yaml:"imageRepository"`
				ImageTag        string `yaml:"imageTag"`
			} `yaml:"local"`
		} `yaml:"etcd"`
		DNS struct {
			Type            string `yaml:"type"`
			ImageRepository string `yaml:"imageRepository"`
			ImageTag        string `yaml:"imageTag"`
		} `yaml:"dns"`
	} `yaml:"kubeadmConfigSpec"`
	Ova []struct {
		Name   string `yaml:"name"`
		Osinfo struct {
			Name    string `yaml:"name"`
			Version string `yaml:"version"`
			Arch    string `yaml:"arch"`
		} `yaml:"osinfo"`
		Version string `yaml:"version"`
	} `yaml:"ova"`
	Ami struct {
		ApNortheast1 []struct {
			ID     string `yaml:"id"`
			Osinfo struct {
				Name    string `yaml:"name"`
				Version string `yaml:"version"`
				Arch    string `yaml:"arch"`
			} `yaml:"osinfo"`
		} `yaml:"ap-northeast-1"`
		ApNortheast2 []struct {
			ID     string `yaml:"id"`
			Osinfo struct {
				Name    string `yaml:"name"`
				Version string `yaml:"version"`
				Arch    string `yaml:"arch"`
			} `yaml:"osinfo"`
		} `yaml:"ap-northeast-2"`
		ApSouth1 []struct {
			ID     string `yaml:"id"`
			Osinfo struct {
				Name    string `yaml:"name"`
				Version string `yaml:"version"`
				Arch    string `yaml:"arch"`
			} `yaml:"osinfo"`
		} `yaml:"ap-south-1"`
		ApSoutheast1 []struct {
			ID     string `yaml:"id"`
			Osinfo struct {
				Name    string `yaml:"name"`
				Version string `yaml:"version"`
				Arch    string `yaml:"arch"`
			} `yaml:"osinfo"`
		} `yaml:"ap-southeast-1"`
		ApSoutheast2 []struct {
			ID     string `yaml:"id"`
			Osinfo struct {
				Name    string `yaml:"name"`
				Version string `yaml:"version"`
				Arch    string `yaml:"arch"`
			} `yaml:"osinfo"`
		} `yaml:"ap-southeast-2"`
		EuCentral1 []struct {
			ID     string `yaml:"id"`
			Osinfo struct {
				Name    string `yaml:"name"`
				Version string `yaml:"version"`
				Arch    string `yaml:"arch"`
			} `yaml:"osinfo"`
		} `yaml:"eu-central-1"`
		EuWest1 []struct {
			ID     string `yaml:"id"`
			Osinfo struct {
				Name    string `yaml:"name"`
				Version string `yaml:"version"`
				Arch    string `yaml:"arch"`
			} `yaml:"osinfo"`
		} `yaml:"eu-west-1"`
		EuWest2 []struct {
			ID     string `yaml:"id"`
			Osinfo struct {
				Name    string `yaml:"name"`
				Version string `yaml:"version"`
				Arch    string `yaml:"arch"`
			} `yaml:"osinfo"`
		} `yaml:"eu-west-2"`
		EuWest3 []struct {
			ID     string `yaml:"id"`
			Osinfo struct {
				Name    string `yaml:"name"`
				Version string `yaml:"version"`
				Arch    string `yaml:"arch"`
			} `yaml:"osinfo"`
		} `yaml:"eu-west-3"`
		SaEast1 []struct {
			ID     string `yaml:"id"`
			Osinfo struct {
				Name    string `yaml:"name"`
				Version string `yaml:"version"`
				Arch    string `yaml:"arch"`
			} `yaml:"osinfo"`
		} `yaml:"sa-east-1"`
		UsEast1 []struct {
			ID     string `yaml:"id"`
			Osinfo struct {
				Name    string `yaml:"name"`
				Version string `yaml:"version"`
				Arch    string `yaml:"arch"`
			} `yaml:"osinfo"`
		} `yaml:"us-east-1"`
		UsEast2 []struct {
			ID     string `yaml:"id"`
			Osinfo struct {
				Name    string `yaml:"name"`
				Version string `yaml:"version"`
				Arch    string `yaml:"arch"`
			} `yaml:"osinfo"`
		} `yaml:"us-east-2"`
		UsGovEast1 []struct {
			ID     string `yaml:"id"`
			Osinfo struct {
				Name    string `yaml:"name"`
				Version string `yaml:"version"`
				Arch    string `yaml:"arch"`
			} `yaml:"osinfo"`
		} `yaml:"us-gov-east-1"`
		UsGovWest1 []struct {
			ID     string `yaml:"id"`
			Osinfo struct {
				Name    string `yaml:"name"`
				Version string `yaml:"version"`
				Arch    string `yaml:"arch"`
			} `yaml:"osinfo"`
		} `yaml:"us-gov-west-1"`
		UsWest2 []struct {
			ID     string `yaml:"id"`
			Osinfo struct {
				Name    string `yaml:"name"`
				Version string `yaml:"version"`
				Arch    string `yaml:"arch"`
			} `yaml:"osinfo"`
		} `yaml:"us-west-2"`
	} `yaml:"ami"`
	Azure []struct {
		Sku             string `yaml:"sku"`
		Publisher       string `yaml:"publisher"`
		Offer           string `yaml:"offer"`
		Version         string `yaml:"version"`
		ThirdPartyImage bool   `yaml:"thirdPartyImage"`
		Osinfo          struct {
			Name    string `yaml:"name"`
			Version string `yaml:"version"`
			Arch    string `yaml:"arch"`
		} `yaml:"osinfo"`
	} `yaml:"azure"`
	ImageConfig struct {
		ImageRepository string `yaml:"imageRepository"`
	} `yaml:"imageConfig"`
	Addons struct {
		AkoOperator struct {
			Category     string   `yaml:"category"`
			ClusterTypes []string `yaml:"clusterTypes"`
			PackageName  string   `yaml:"packageName"`
			Reposoitory  string   `yaml:"repository"`
		} `yaml:"ako-operator"`
		Antrea struct {
			Category     string   `yaml:"category"`
			ClusterTypes []string `yaml:"clusterTypes"`
			PackageName  string   `yaml:"packageName"`
			Reposoitory  string   `yaml:"repository"`
		} `yaml:"antrea"`
		Calico struct {
			Category     string   `yaml:"category"`
			ClusterTypes []string `yaml:"clusterTypes"`
			PackageName  string   `yaml:"packageName"`
			Reposoitory  string   `yaml:"repository"`
		} `yaml:"calico"`
		KappController struct {
			Category     string   `yaml:"category"`
			ClusterTypes []string `yaml:"clusterTypes"`
			PackageName  string   `yaml:"packageName"`
			Reposoitory  string   `yaml:"repository"`
		} `yaml:"kapp-controller"`
		LoadBalancerAndIngressService struct {
			Category     string   `yaml:"category"`
			ClusterTypes []string `yaml:"clusterTypes"`
			PackageName  string   `yaml:"packageName"`
			Reposoitory  string   `yaml:"repository"`
		} `yaml:"load-balancer-and-ingress-service"`
		MetricsServer struct {
			Category     string   `yaml:"category"`
			ClusterTypes []string `yaml:"clusterTypes"`
			PackageName  string   `yaml:"packageName"`
			Reposoitory  string   `yaml:"repository"`
		} `yaml:"metrics-server"`
		Pinniped struct {
			Category     string   `yaml:"category"`
			ClusterTypes []string `yaml:"clusterTypes"`
			PackageName  string   `yaml:"packageName"`
			Reposoitory  string   `yaml:"repository"`
		} `yaml:"pinniped"`
		TanzuAddonsManager struct {
			Category     string   `yaml:"category"`
			ClusterTypes []string `yaml:"clusterTypes"`
			PackageName  string   `yaml:"packageName"`
			Reposoitory  string   `yaml:"repository"`
		} `yaml:"tanzu-addons-manager"`
		VsphereCpi struct {
			Category     string   `yaml:"category"`
			ClusterTypes []string `yaml:"clusterTypes"`
			PackageName  string   `yaml:"packageName"`
			Reposoitory  string   `yaml:"repository"`
		} `yaml:"vsphere-cpi"`
		VsphereCsi struct {
			Category     string   `yaml:"category"`
			ClusterTypes []string `yaml:"clusterTypes"`
			PackageName  string   `yaml:"packageName"`
			Reposoitory  string   `yaml:"repository"`
		} `yaml:"vsphere-csi"`
	} `yaml:"addons"`
}

func ReadTKRBom(filePath string) (*TKRBom, error) {
	bom := &TKRBom{}
	rawBom, err := os.ReadFile(filePath)
	if err != nil {
		return bom, err
	}
	err = yaml.Unmarshal(rawBom, bom)
	if err != nil {
		return bom, err
	}

	return bom, nil
}

func (tkr *TKRBom) getTKRRegistry() string {
	return tkr.ImageConfig.ImageRepository
}

func (tkr *TKRBom) GetTKRNodeImage() string {
	repo := tkr.getTKRNodeRepository()
	path := tkr.Components.KubernetesSigsKind[0].Images.KindNodeImage.ImagePath
	tag := tkr.Components.KubernetesSigsKind[0].Images.KindNodeImage.Tag

	return fmt.Sprintf("%s/%s:%s", repo, path, tag)
}

func (tkr *TKRBom) GetTKRCoreRepoBundlePath() string {
	registry := tkr.getTKRRegistry()
	path := tkr.Components.TkgCorePackages[0].Images.TanzuCorePackageRepositoryImage.ImagePath
	tag := tkr.Components.TkgCorePackages[0].Images.TanzuCorePackageRepositoryImage.Tag

	return fmt.Sprintf("%s/%s:%s", registry, path, tag)
}

// TODO(joshrosso): We're waiting on this information to be available in the TKR API
// for now, we are hard-coding the response
func (tkr *TKRBom) GetAdditionalRepoBundlesPaths() []string {
	return []string{tceRepoURL}
}

func (tkr *TKRBom) GetTKRKappImage() (TkrImageReader, error) {
	registry := tkr.getTKRKappRepository()
	path := tkr.Components.TkgCorePackages[0].Images.KappControllerTanzuVmwareCom.ImagePath
	tag := tkr.Components.TkgCorePackages[0].Images.KappControllerTanzuVmwareCom.Tag

	t, err := NewTkrImageReader(fmt.Sprintf("%s/%s:%s", registry, path, tag))
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (tkr *TKRBom) getTKRNodeRepository() string {
	if tkr.Components.KubernetesSigsKind[0].Images.KindNodeImage.Reposoitory == "" {
		return tkr.getTKRRegistry()
	}

	return tkr.Components.KubernetesSigsKind[0].Images.KindNodeImage.Reposoitory
}

func (tkr *TKRBom) getTKRKappRepository() string {
	if tkr.Components.TkgCorePackages[0].Images.KappControllerTanzuVmwareCom.Reposoitory == "" {
		return tkr.getTKRRegistry()
	}

	return tkr.Components.TkgCorePackages[0].Images.KappControllerTanzuVmwareCom.Reposoitory
}