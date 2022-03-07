import re
import sys

# regex:
MODULE_REGEX = r'^[_a-zA-Z][_a-zA-Z0-9]+$'


def validate_name(regex):
    # name list:
    names = [
        '{{cookiecutter.repo_name}}',
        '{{cookiecutter.biz_app_name}}',
    ]

    # validate:
    for name in names:
        if not re.match(regex, name):
            print('ERROR: The project name (%s) is not a valid module name. Please do not use a - and use _ instead' % name)
            return False
    # ok:
    return True


if __name__ == '__main__':
    # validate:
    if not validate_name(MODULE_REGEX):
        # Exit to cancel project
        sys.exit(1)
