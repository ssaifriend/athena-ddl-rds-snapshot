package ddl

const DdlTemplate  = `
{{define "DDL" -}}
CREATE EXTERNAL TABLE IF NOT EXISTS ` + "`{{.AthenaDB}}.{{.Table.AthenaName}}`" + `(
  {{ range $i, $e := .Table.Columns -}}
` + "`{{$e.Name}}`" + ` {{$e.AthenaType}}{{if last $i $.Table.Columns | not }},{{end}}
  {{ end -}}
)
ROW FORMAT SERDE 'org.apache.hadoop.hive.ql.io.parquet.serde.ParquetHiveSerDe'
WITH SERDEPROPERTIES (
  'serialization.format' = '1'
) LOCATION '{{.S3Prefix}}public.{{.Table.AthenaName}}/'
TBLPROPERTIES ('has_encrypted_data'='false');
{{end}}
`