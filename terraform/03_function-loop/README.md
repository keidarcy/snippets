## references

- [count](https://developer.hashicorp.com/terraform/language/meta-arguments/count)
- [for_each](https://developer.hashicorp.com/terraform/language/meta-arguments/for_each)
- [functions](https://developer.hashicorp.com/terraform/language/functions)
- [expressions](https://developer.hashicorp.com/terraform/language/expressions)

## `terraform console`

### function

```
> min(1,23,4)
1
> max(1,23,4)
23
> lower("TEsT")
"test"
> upper("a")
"A"
> var.vpc_cidr_block
tomap({
  "private" = "10.1.0.0/16"
  "public" = "192.168.0.0/16"
})
> lookup(var.vpc_cidr_block, "public")
"192.168.0.0/16"
> lookup(var.vpc_cidr_block, "public1")
╷
│ Error: Error in function call
│
│   on <console-input> line 1:
│   (source code not available)
│
│ Call to function "lookup" failed: lookup failed to find key "public1".
╵


> lookup(var.vpc_cidr_block, "public1", "unknown")
"unknown"
> merge(var.vpc_cidr_block, {"public1": "1.1.1.1/32"})
{
  "private" = "10.1.0.0/16"
  "public" = "192.168.0.0/16"
  "public1" = "1.1.1.1/32"
}
> file("${path.module}/test.tpl")
"hello world ${name}"
> templatefile("${path.module}/test.tpl", {"name": "terraform name"})
"hello world terraform name"
```

### expression

```
> var.aws_regions
tolist([
  "eu-central-1",
  "us-east-1",
  "us-east-2",
])

> [for v in var.aws_regions: upper(v)]
[
  "EU-CENTRAL-1",
  "US-EAST-1",
  "US-EAST-2",
]

> [for i, v in var.aws_regions: "${i} upper is ${upper(v)}"]
[
  "0 upper is EU-CENTRAL-1",
  "1 upper is US-EAST-1",
  "2 upper is US-EAST-2",
]

> { for i, s in var.aws_regions: "key-${s}" => "${i}-${upper(s)}" }
{
  "key-eu-central-1" = "0-EU-CENTRAL-1"
  "key-us-east-1" = "1-US-EAST-1"
  "key-us-east-2" = "2-US-EAST-2"
}

```