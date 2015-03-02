try:
    from setuptools import setup, Extension
except ImportError:
    from distutils.core import setup, Extension

import golang_ext

gotypes_extension = Extension('_gotypes', sources=['_gotypes.go'])


__author__ = 'Jon Nappi'

with open('README.rst') as f:
    readme = f.read()
with open('HISTORY.rst') as f:
    history = f.read()
with open('requirements.txt') as f:
    requires = [line.strip() for line in f if line.strip()]

packages = ['golang_ext', 'gotypes']
description = 'Python Golang Extension Builder and gotypes bindings.'

setup(
    name='golang_ext',
    version=golang_ext.__version__,
    description=description,
    long_description='\n'.join([readme, history]),
    author='Jonathan Nappi',
    author_email='moogar0880@gmail.com',
    url='https://github.com/moogar0880/golang_ext',
    packages=packages,
    install_requires=requires,
    ext_modules=[gotypes_extension],
    cmdclass={'build_ext': golang_ext.GolangBuildExt},
    license='Apache 2.0',
    zip_safe=False,
    classifiers=(
        'Development Status :: 1 - Planning',
        'Intended Audience :: Developers',
        'Natural Language :: English',
        'License :: Freely Distributable',
        'Programming Language :: Python',
        'Programming Language :: Python :: 2.7',
        'Programming Language :: Python :: 3',
        'Programming Language :: Python :: 3.4')
)
