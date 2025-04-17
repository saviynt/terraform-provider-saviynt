resource "saviynt_github_rest_connection_resource" "example3" {
  connection_type     = "GithubRest"
  connection_name     = "namefortheconnection"
  connection_json=jsonencode({
      "objects":
          "objectClasses": [
            "user",
            "top",
            "Person",
            "OrganizationalPerson"
          ],
    })
  import_account_ent_json=jsonencode({
      "objects":
          "objectClasses": [
            "user",
            "top",
            "Person",
            "OrganizationalPerson"
          ],
    })
  access_tokens="XXXXX"
  organization_list="example"
  status_threshold_config=jsonencode({
      "objects":
          "objectClasses": [
            "user",
            "top",
            "Person",
            "OrganizationalPerson"
          ],
    })
  pam_config=jsonencode({
      "objects":
          "objectClasses": [
            "user",
            "top",
            "Person",
            "OrganizationalPerson"
          ],
    })
}