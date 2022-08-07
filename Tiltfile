# -*- mode: Python -*-

# For more on Extensions, see: https://docs.tilt.dev/extensions.html
load('ext://restart_process', 'docker_build_with_restart')

compile_cmd = 'CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -gcflags "-N -l" -o build/tilt-rn-agent ./'
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
  # entrypoint=['/app/build/tilt-rn-agent'],
  entrypoint='dlv debug --headless --log -l 0.0.0.0:2345 --api-version=2 --accept-multiclient',
  dockerfile='deployments/Dockerfile',
  live_update=[
    sync('./build', '/app/build'),
    sync('./web', '/app/web'),
    sync('./crt', '/app/crt'),
  ],
)

k8s_yaml('deployments/kubernetes.yaml')
k8s_resource('rn-agent', port_forwards=[8000, 2345],
             resource_deps=['rn-agent-compile'])
