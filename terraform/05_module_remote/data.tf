data "template_file" "cloud-init" {
  template = file("${path.module}/data/cloud-init.yaml")
}
