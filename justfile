[private]
default:
  just --list

publish repository image tag arch: (insert image tag arch)
  skopeo --insecure-policy copy oci://$(pwd)/{{image}} docker://{{repository}}/{{image}}:{{tag}}

[private]
insert image tag arch: (binary image tag arch) && (garbage image)
  umoci insert \
    --image {{image}}:{{tag}} \
    --opaque \
    app /bin/app

[private]
binary image tag arch: (config image tag arch)
  CGO_ENABLED=0 GOOS=linux GOARCH={{arch}} \
    go build -o app

[private]
config image tag arch: (new image tag)
  umoci config \
    --image {{image}} \
    --os linux \
    --architecture {{arch}} \
    --config.entrypoint /bin/app

[private]
new image tag: (init image)
  umoci new --image {{image}}:{{tag}}

[private]
init image:
  umoci init --layout {{image}}

[private]
garbage image:
  umoci gc --layout {{image}}

[private]
clean image:
  rm -rf {{image}}
