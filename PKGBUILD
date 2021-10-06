# Maintainer: Taylor McKinnon <mail@tmacs.space>
pkgname='today-dir-git'
pkgver='0'
pkgrel=1
pkgdesc='Simple scratch directory and notes manager'
arch=('any')
url="https://github.com/tmacro/today"
license=('BSD-3-Clause')
groups=()
depends=()
makedepends=('git' 'python-setuptools')
provides=("${pkgname%-git}")
conflicts=("${pkgname%-git	}")
replaces=()
backup=()
options=()
install=
source=('git+https://github.com/tmacro/today.git')
noextract=()
md5sums=('SKIP')

pkgver() {
	cd "$srcdir/${pkgname%-dir-git}"
	printf "r%s.%s" "$(git rev-list --count HEAD)" "$(git rev-parse --short HEAD)"
}

build() {
	cd "$srcdir/${pkgname%-dir-git}"
	python setup.py build
}

package() {
	cd "$srcdir/${pkgname%-dir-git}"
	python setup.py install --root="${pkgdir}" --optimize=1 --skip-build
}
