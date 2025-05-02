terraform {
  required_providers {
    saviynt = {
      source  = "registry.terraform.io/local/saviynt"
      version = "1.0.0"
    }
  }
}

provider "saviynt" {
  server_url = var.SAVIYNT_SERVER_URL
  username   = var.SAVIYNT_USERNAME
  password   = var.SAVIYNT_PASSWORD
}

resource "saviynt_unix_connection_resource" "example" {
  connection_type = "Unix"
  connection_name = "Terraform_Unix_Connector"
  host_name       = "unix.example.com"
  port_number     = "22"
  username        = "provision_user"
  password        = "SuperSecretPassword123!"

  groups_file   = "/etc/group"
  accounts_file = "/etc/passwd"
  shadow_file   = jsonencode({ UNIX = "sudo cat /etc/shadow" })

  provision_account_command = "sudo useradd $$${user.username} -p $$${user.password} -c \"$$${user.fullname}\" -g users"

  deprovision_account_command = "sudo userdel -r $$${user.username}"

  add_access_command = "sudo usermod -aG $$${entitlement.groupname} $$${user.username}"

  remove_access_command = "sudo gpasswd -d $$${user.username} $$${entitlement.groupname}"

  change_password_json = jsonencode({
    UNIX = "echo $$${user.username}:$$${user.new_password} | sudo chpasswd"
  })

  pem_key_file = "/etc/saviynt/keys/id_rsa"

  enable_account_command = "sudo passwd -u $$${user.username}"

  disable_account_command = "sudo passwd -l $$${user.username}"

  account_entitlement_mapping_command = "grep $$${user.username} /etc/group"

  passphrase = "optional_passphrase"

  update_account_command = "sudo usermod -c \"$$${user.comment}\" $$${user.username}"

  create_group_command = "sudo groupadd $$${group.name}"

  delete_group_command = "sudo groupdel $$${group.name}"

  add_group_owner_command = "sudo chown $$${group.owner} /etc/group"

  add_primary_group_command = "sudo usermod -g $$${group.name} $$${user.username}"

  fire_fighter_id_grant_access_command = "sudo usermod -aG firegroup $$${user.username}"

  fire_fighter_id_revoke_access_command = "sudo gpasswd -d $$${user.username} firegroup"

  inactive_lock_account = "true"

  status_threshold_config = jsonencode({
    threshold  = 3
    time_unit  = "days"
    grace_days = 2
  })

  custom_config_json = jsonencode({
    shell_type = "bash"
    timeout    = 30
    retryPolicy = {
      maxAttempts = 3
      backoff     = "linear"
    }
  })

  ssh_key = <<EOKEY
  -----BEGIN RSA PRIVATE KEY-----
  MIICXQIBAAKBgQDX...
  -----END RSA PRIVATE KEY-----
  EOKEY

  lock_account_command   = "sudo usermod -L $$${user.username}"
  unlock_account_command = "sudo usermod -U $$${user.username}"

  pass_through_connection_details = jsonencode({
    hostname = "unix-jumphost.example.com"
    port     = "22"
    username = "jumphost_user"
  })

  ssh_pass_through_password   = "JumpHostSecretPass#456"
  ssh_pass_through_sshkey     = "/home/jumphost_user/.ssh/id_rsa"
  ssh_pass_through_passphrase = "JumphostKeyPassphrase"
}