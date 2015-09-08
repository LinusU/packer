package ansible

import (
	"fmt"
	"os"

	"github.com/mitchellh/packer/common"
	"github.com/mitchellh/packer/helper/config"
	"github.com/mitchellh/packer/packer"
	"github.com/mitchellh/packer/template/interpolate"
)

type Config struct {
	common.PackerConfig `mapstructure:",squash"`
	ctx                 interpolate.Context

	// The playbook dir to upload.
	PlaybookDir string `mapstructure:"playbook_dir"`

	// The main playbook file to execute.
	PlaybookFile string `mapstructure:"playbook_file"`
}

type Provisioner struct {
	config Config
}
func (p *Provisioner) Prepare(raws ...interface{}) error {
	err := config.Decode(&p.config, &config.DecodeOpts{
		Interpolate:        true,
		InterpolateContext: &p.config.ctx,
		InterpolateFilter: &interpolate.RenderFilter{
			Exclude: []string{},
		},
	}, raws...)
	if err != nil {
		return err
	}

	// Validation
	var errs *packer.MultiError
	err = validateFileConfig(p.config.PlaybookFile, "playbook_file", true)
	if err != nil {
		errs = packer.MultiErrorAppend(errs, err)
	}

	// Check that the playbook_dir directory exists, if configured
	if len(p.config.PlaybookDir) > 0 {
		if err := validateDirConfig(p.config.PlaybookDir, "playbook_dir"); err != nil {
			errs = packer.MultiErrorAppend(errs, err)
		}
	}

	if errs != nil && len(errs.Errors) > 0 {
		return errs
	}
	return nil
}


func (p *Provisioner) Provision(ui packer.Ui, comm packer.Communicator) error {
	ui.Say("Provisioning with Ansible...")

  var dasdasd = fmt.Sprintf("%s ansible_ssh_host=%s ansible_ssh_port=%d ansible_ssh_user=%s ansible_ssh_private_key_file=%s", "test-server-1", comm.config.SSHHost, comm.config.SSHPort, comm.config.SSHUsername, comm.config.SSHPrivateKey)

  ui.Say(dasdasd)













	// if len(p.config.PlaybookDir) > 0 {
	// 	ui.Message("Uploading Playbook directory to Ansible staging directory...")
	// 	if err := p.uploadDir(ui, comm, p.config.StagingDir, p.config.PlaybookDir); err != nil {
	// 		return fmt.Errorf("Error uploading playbook_dir directory: %s", err)
	// 	}
	// } else {
	// 	ui.Message("Creating Ansible staging directory...")
	// 	if err := p.createDir(ui, comm, p.config.StagingDir); err != nil {
	// 		return fmt.Errorf("Error creating staging directory: %s", err)
	// 	}
	// }
  //
	// ui.Message("Uploading main Playbook file...")
	// src := p.config.PlaybookFile
	// dst := filepath.ToSlash(filepath.Join(p.config.StagingDir, filepath.Base(src)))
	// if err := p.uploadFile(ui, comm, dst, src); err != nil {
	// 	return fmt.Errorf("Error uploading main playbook: %s", err)
	// }
  //
	// if len(p.config.InventoryFile) == 0 {
	// 	tf, err := ioutil.TempFile("", "packer-provisioner-ansible-local")
	// 	if err != nil {
	// 		return fmt.Errorf("Error preparing inventory file: %s", err)
	// 	}
	// 	defer os.Remove(tf.Name())
	// 	if len(p.config.InventoryGroups) != 0 {
	// 		content := ""
	// 		for _, group := range p.config.InventoryGroups {
	// 			content += fmt.Sprintf("[%s]\n127.0.0.1\n", group)
	// 		}
	// 		_, err = tf.Write([]byte(content))
	// 	} else {
	// 		_, err = tf.Write([]byte("127.0.0.1"))
	// 	}
	// 	if err != nil {
	// 		tf.Close()
	// 		return fmt.Errorf("Error preparing inventory file: %s", err)
	// 	}
	// 	tf.Close()
	// 	p.config.InventoryFile = tf.Name()
	// 	defer func() {
	// 		p.config.InventoryFile = ""
	// 	}()
	// }
  //
	// ui.Message("Uploading inventory file...")
	// src = p.config.InventoryFile
	// dst = filepath.ToSlash(filepath.Join(p.config.StagingDir, filepath.Base(src)))
	// if err := p.uploadFile(ui, comm, dst, src); err != nil {
	// 	return fmt.Errorf("Error uploading inventory file: %s", err)
	// }
  //
	// if len(p.config.GroupVars) > 0 {
	// 	ui.Message("Uploading group_vars directory...")
	// 	src := p.config.GroupVars
	// 	dst := filepath.ToSlash(filepath.Join(p.config.StagingDir, "group_vars"))
	// 	if err := p.uploadDir(ui, comm, dst, src); err != nil {
	// 		return fmt.Errorf("Error uploading group_vars directory: %s", err)
	// 	}
	// }
  //
	// if len(p.config.HostVars) > 0 {
	// 	ui.Message("Uploading host_vars directory...")
	// 	src := p.config.HostVars
	// 	dst := filepath.ToSlash(filepath.Join(p.config.StagingDir, "host_vars"))
	// 	if err := p.uploadDir(ui, comm, dst, src); err != nil {
	// 		return fmt.Errorf("Error uploading host_vars directory: %s", err)
	// 	}
	// }
  //
	// if len(p.config.RolePaths) > 0 {
	// 	ui.Message("Uploading role directories...")
	// 	for _, src := range p.config.RolePaths {
	// 		dst := filepath.ToSlash(filepath.Join(p.config.StagingDir, "roles", filepath.Base(src)))
	// 		if err := p.uploadDir(ui, comm, dst, src); err != nil {
	// 			return fmt.Errorf("Error uploading roles: %s", err)
	// 		}
	// 	}
	// }
  //
	// if len(p.config.PlaybookPaths) > 0 {
	// 	ui.Message("Uploading additional Playbooks...")
	// 	playbookDir := filepath.ToSlash(filepath.Join(p.config.StagingDir, "playbooks"))
	// 	if err := p.createDir(ui, comm, playbookDir); err != nil {
	// 		return fmt.Errorf("Error creating playbooks directory: %s", err)
	// 	}
	// 	for _, src := range p.config.PlaybookPaths {
	// 		dst := filepath.ToSlash(filepath.Join(playbookDir, filepath.Base(src)))
	// 		if err := p.uploadDir(ui, comm, dst, src); err != nil {
	// 			return fmt.Errorf("Error uploading playbooks: %s", err)
	// 		}
	// 	}
	// }
  //
	// if err := p.executeAnsible(ui, comm); err != nil {
	// 	return fmt.Errorf("Error executing Ansible: %s", err)
	// }
	return nil
}

func (p *Provisioner) Cancel() {
	// Just hard quit. It isn't a big deal if what we're doing keeps
	// running on the other side.
	os.Exit(0)
}


















func validateDirConfig(path string, config string) error {
	info, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("%s: %s is invalid: %s", config, path, err)
	} else if !info.IsDir() {
		return fmt.Errorf("%s: %s must point to a directory", config, path)
	}
	return nil
}

func validateFileConfig(name string, config string, req bool) error {
	if req {
		if name == "" {
			return fmt.Errorf("%s must be specified.", config)
		}
	}
	info, err := os.Stat(name)
	if err != nil {
		return fmt.Errorf("%s: %s is invalid: %s", config, name, err)
	} else if info.IsDir() {
		return fmt.Errorf("%s: %s must point to a file", config, name)
	}
	return nil
}
