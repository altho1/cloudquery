# Table: azure_sql_database_threat_detection_policies

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/v4.0/sql#DatabaseSecurityAlertPolicy

The primary key for this table is **id**.

## Relations
This table depends on [azure_sql_databases](azure_sql_databases.md).

## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|sql_database_id|String|
|location|String|
|kind|String|
|state|String|
|disabled_alerts|String|
|email_addresses|String|
|email_account_admins|String|
|storage_endpoint|String|
|storage_account_access_key|String|
|retention_days|Int|
|use_server_default|String|
|id (PK)|String|
|name|String|
|type|String|