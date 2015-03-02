# -*- coding: utf-8 -*-
from setuptools.command import build_ext
from distutils.unixccompiler import UnixCCompiler

__author__ = 'Jon Nappi'
__all__ = ['GolangBuildExt', 'GolangCompiler']


class GolangBuildExt(build_ext.build_ext):
    """Custom extension builder that sets builds extensions using the
    :class:`GolangCompiler` as the compiler class
    """

    def initialize_options(self):
        super(GolangBuildExt, self).initialize_options()
        self.compiler = GolangCompiler


class GolangCompiler(UnixCCompiler):
    """A Custom Compiler class type detailing how to use gccgo to compile go
    code into a C binary that Python can consume
    """

    compiler_type = 'go'

    # These are used by CCompiler in two places: the constructor sets
    # instance attributes 'preprocessor', 'compiler', etc. from them, and
    # 'set_executable()' allows any of these to be set.  The defaults here
    # are pretty generic; they will probably have to be set by an outsider
    # (eg. using information discovered by the sysconfig about building
    # Python extensions).
    executables = {
        'preprocessor': None,
        'compiler': ['gccgo'],
        'compiler_so': ['gccgo'],
        'compiler_cxx': None,
        'linker_so': ['gccgo'],
        'linker_exe': None,
        'archiver': ['ar', '-cr'],
        'ranlib': None,
    }

    # `language_map` is used to detect a source file or Extension target
    # language, checking source filenames. language_order is used to detect
    # the language precedence, when deciding what language to use when mixing
    # source types. For example, if some extension has two files with ".c"
    # extension, and one with ".cpp", it is still linked as c++. However, given
    # that this Compiler is only concerned with compiling .go files we don't
    # need to worry about compilation precendence
    language_map = {'.go': 'go'}
    language_order = ['go']
    src_extensions = ['.go']

    def __init__(self, verbose=0, dry_run=0, force=0):
        """Custom compiler init override implicitly sets the required base
        gccgo include directories
        """
        super(GolangCompiler, self).__init__(verbose, dry_run, force)
        self.include_dirs = ['/usr/lib/gccgo', '/usr/local/lib/gccgo']

    def library_dir_option(self, dir):
        """Return the compiler option to add 'dir' to the list of libraries
        linked into the shared library or executable.
        """
        return ''.join(['-L', dir])

    def library_option(self, lib):
        """Return the compiler option to add 'dir' to the list of libraries
        linked into the shared library or executable.
        """
        return ''.join(['-l', dir])
