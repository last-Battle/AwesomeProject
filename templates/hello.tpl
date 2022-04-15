<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Hello</title>
</head>
<body>
    $obj := {{.}}
    <p>Hello {{- .Name -}}</p>
    <p>性别：{{.Gender}}</p>
    <p>年龄：{{.Age}}</p>

    {{if pipeline }} T1 {{end}}
    {{if pipeline }} T1 {{else}} T0 {{end}}

    {{range pipeline }}T1 {{end}}
    {{eq arg1 arg2 arg3}}

    {{ /* define "ol.tmpl" */}}
    {{end}}

</body>
</html>