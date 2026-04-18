#!/usr/bin/env sh
set -eu

script_dir=$(CDPATH= cd -- "$(dirname "$0")" && pwd)
binary_path="$script_dir/bin/update-gowork"
source_path="$script_dir/update-gowork.go"

cd "$script_dir"

usage() {
	cat <<EOF
usage: ./update-gowork.sh [--build-binary [output]] [--run-go] [--clean] [--help]

Synchronize go.work with the modules found in the project tree.
The default mode runs the compiled binary from bin/update-gowork and
rebuilds it automatically when update-gowork.go is newer or missing.

Options:
  --build-binary [output]  Build the synchronizer binary.
                           Default output: $binary_path
  --run-go                 Run the Go source directly with GOWORK=off.
                           Useful when you want to test the source itself.
  --clean                  Remove the generated binary from:
                           $binary_path
  --help, -h               Show this help message.

Examples:
  ./update-gowork.sh
  ./update-gowork.sh --build-binary
  ./update-gowork.sh --build-binary ./bin/custom-update-gowork
  ./update-gowork.sh --run-go
  ./update-gowork.sh --clean

Notes:
  - The Go source is always executed with GOWORK=off when using --run-go.
  - The binary is rebuilt only when needed in the default mode.
  - The compiled binary can be run directly from any location with:
    $binary_path --dir /path/to/project
EOF
}

build_binary() {
	output_path=${1:-$binary_path}
	mkdir -p "$(dirname "$output_path")"
	GOWORK=off go build -o "$output_path" ./update-gowork.go
	echo "Built $output_path"
}

clean_binary() {
	if [ -f "$binary_path" ]; then
		rm -f "$binary_path"
		echo "Removed $binary_path"
		return
	fi

	echo "No binary to remove"
}

run_binary() {
	if [ ! -x "$binary_path" ] || [ "$source_path" -nt "$binary_path" ]; then
		build_binary "$binary_path"
	fi

	"$binary_path" --dir "$script_dir"
}

case "${1:-}" in
	"" )
		run_binary
		;;
	--build-binary)
		build_binary "${2:-}"
		;;
	--run-go)
		GOWORK=off go run ./update-gowork.go
		;;
	--clean)
		clean_binary
		;;
	--help|-h)
		usage
		;;
	*)
		echo "invalid option: $1" >&2
		usage >&2
		exit 1
		;;
	esac