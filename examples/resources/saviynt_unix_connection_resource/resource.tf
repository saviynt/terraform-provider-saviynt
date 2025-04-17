
resource "saviynt_unix_connection_resource" "example" {
  connection_type = "Unix"
  connection_name = "namefortheconnection"
  host_name                      = "unix.example.com"
  port_number                    = "22"
  username                       = "provisionuser"
  password                       = "MyUnix@123"
  groups_file                    = "{\"UNIX\":\"sudo cat /etc\"}"
  accounts_file                  = "{\"UNIX\":\"sudo cat /etc\"}"
  shadow_file                    = "{\"UNIX\":\"sudo cat /etc\"}"
  provision_account_command      = "sudo useradd $${username} -p $${password} -g users"
  deprovision_account_command    = "sudo userdel $${username}"
  add_access_command             = "sudo usermod -a -G $${group} $${username}"
  remove_access_command          = "sudo gpasswd -d $${username} $${group}"
  change_password_json           = "{\"UNIX\":{\"cmd\":\"echo $${username}:$${newPassword} | chpasswd\"}}"
  pem_key_file                   = "/etc/saviynt/keys/id_rsa"
  enable_account_command         = "sudo usermod -U $${username}"
  disable_account_command        = "sudo usermod -L $${username}"
  account_entitlement_mapping_command = "sudo cat /etc/group | grep $${username}"
  passphrase                     = "MySecurePassphrase"
  update_account_command         = "sudo usermod -c \"$${email}\" $${username}"
  create_group_command           = "sudo groupadd $${groupname}"
  delete_group_command           = "sudo groupdel $${groupname}"
  add_group_owner_command        = "sudo chown $${owner} /etc/group"
  add_primary_group_command      = "sudo usermod -g $${primaryGroup} $${username}"
  fire_fighter_id_grant_access_command = "sudo usermod -a -G emergency $${username}"
  fire_fighter_id_revoke_access_command = "sudo gpasswd -d $${username} emergency"
  inactive_lock_account          = "sudo usermod -L $${username}"

  status_threshold_config = <<EOT
  {
    "statusAndThresholdConfig": {
      "accountStatusAttribute": "status",
      "activeStatusValue": ["ACTIVE"],
      "disableStatusValue": ["LOCKED"],
      "thresholdKey": "lastLogin",
      "lockoutValue": "30d"
    }
  }
  EOT

  custom_config_json             = "{\"maxRetries\":3,\"timeout\":10}"
  ssh_key                        = "-----BEGIN RSA PRIVATE KEY-----\nMIIBOgIBAAJBAK...\n-----END RSA PRIVATE KEY-----"
  lock_account_command           = "sudo usermod -L $${username}"
  unlock_account_command         = "sudo usermod -U $${username}"
  pass_through_connection_details = "{\"host\":\"bastion.example.com\",\"port\":22}"
  ssh_pass_through_password      = "XXXXX"
  ssh_pass_through_sshkey        = "-----BEGIN OPENSSH PRIVATE KEY-----\n..."
  ssh_pass_through_passphrase    = "BastionPassphrase"

}
   