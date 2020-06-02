#!/usr/bin/env python
import os

PROJECT_DIRECTORY = os.path.realpath(os.path.curdir)


def remove_file(filepath):
    os.remove(os.path.join(PROJECT_DIRECTORY, filepath))


def choice():
    # 选择:
    if '{{ cookiecutter.folder_type }}' == 'Single Micro Service: App Folder':
        single_app_create()  # 创建 单个 app 目录
    else:
        mono_repo_create()  # 创建 mono repo 根目录


def mono_repo_create():
    print("single app create:")


def single_app_create():
    print("single app create:")


if __name__ == '__main__':
    choice()

    # clean:
    if 'Not open source' == '{{ cookiecutter.open_source_license }}':
        remove_file('LICENSE')
