resource "saviynt_dynamic_attribute_resource" "example" {
  security_system = "SAP"
  endpoint        = "SAP-PRD"
  user_name       = "admin"

  dynamic_attribute = {
    attribute_name = "attribute-name" #Compulsory for all actions
    request_type   = "Accounts"       #Compulsory for creation
    update_user = "username" #optional for creation, compulsory for updation and deletion
    # attribute_type     = "BOOLEAN"
    # attribute_group    = "ACCOUNT-GROUP"
    # order_index        = "1"
    # attribute_lable    = "ATTRIBUTE-LABEL"
    # accounts_column    = "ACCOUNTS-COLUMN"
    # hide_on_create     = "false"
    # action_string      = ""
    # editable           = "true"
    # hide_on_update     = "false"
    # actiontoperformwhenparentattributechanges = ""
    # default_value      = "Engineering"
    # required           = "true"
    # regex              = ".*"
    # attribute_value    = ""
    # showonchild        = "false"
    # parentattribute    = ""
    # descriptionascsv   = "Sample description"
  }

}

