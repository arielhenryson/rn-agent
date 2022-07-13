# -*- mode: Python -*-

# For more on Extensions, see: https://docs.tilt.dev/extensions.html
load('ext://restart_process', 'docker_build_with_restart')

compile_cmd = 'CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/tilt-rn-agent ./'
if os.name == 'nt':
  compile_cmd = 'build.bat'

local_resource(
  'rn-agent-compile',
  compile_cmd,
  deps=['./main.go'],
  resource_deps=[])

docker_build_with_restart(
  'rn-agent-image',
  '.',
  entrypoint=['/app/build/tilt-rn-agent'],
  dockerfile='deployments/Dockerfile',
  only=[
    './build',
    './web',
  ],
  live_update=[
    sync('./build', '/app/build'),
    sync('./web', '/app/web'),
  ],
)

k8s_yaml('deployments/kubernetes.yaml')
k8s_resource('rn-agent', port_forwards=8000,
             resource_deps=['rn-agent-compile'])
