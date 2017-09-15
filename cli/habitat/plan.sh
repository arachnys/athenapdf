pkg_name=athenapdf-cli
pkg_origin=jarvus
pkg_version=2.10.0
pkg_bin_dirs=(bin)
pkg_lib_dirs=(lib)

pkg_deps=(
  core/glibc
  core/gcc-libs
  core/coreutils
  core/node

  core/freetype
  core/fontconfig
  core/dbus
  core/expat
  core/cairo
  core/pango
  core/nss
  core/nspr
  core/glib
  xorg/xlib
  xorg/libxrender
  xorg/libxext
  xorg/libxcb
)
pkg_build_deps=(
  core/patchelf
)

do_unpack() {
  pushd "${CACHE_PATH}" > /dev/null

  cp -vr "${PLAN_CONTEXT}/../src" .
  cp -v "${PLAN_CONTEXT}/../package.json" .

  popd > /dev/null

  return 0
}

do_build() {
  pushd "${CACHE_PATH}" > /dev/null

  npm install
  fix_interpreter "node_modules/.bin/*" core/coreutils bin/env
  node_modules/.bin/electron-packager . athenapdf --platform=linux --arch=x64 --version=1.7.5
  patchelf --interpreter "$(pkg_path_for glibc)/lib/ld-linux-x86-64.so.2" --set-rpath "${LD_RUN_PATH}" athenapdf-linux-x64/athenapdf

  find athenapdf-linux-x64 -name "*.so" \
    -exec patchelf --set-rpath "${LD_RUN_PATH}" {} \;

  popd > /dev/null
}

do_install() {
  pushd "${pkg_prefix}" > /dev/null

  cp -v "${CACHE_PATH}/athenapdf-linux-x64"/*.so lib/
  cp -v "${CACHE_PATH}/athenapdf-linux-x64/athenapdf" bin/

  ldd bin/*
  attach

  popd > /dev/null
}