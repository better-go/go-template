#!/usr/bin/env python
import os

PROJECT_DIRECTORY = os.path.realpath(os.path.curdir)


def remove_file(filepath):
    os.remove(os.path.join(PROJECT_DIRECTORY, filepath))


def mono_repo_create():
    print("mono repo create:")


if __name__ == '__main__':
    mono_repo_create()

    # clean:
    if 'Not open source' == '{{ cookiecutter.open_source_license }}':
        remove_file('LICENSE')
