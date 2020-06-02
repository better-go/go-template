import re
import sys

# regex:
MODULE_REGEX = r'^[_a-zA-Z][_a-zA-Z0-9]+$'


def validate_name(regex):
    # name list:
    names = [
        '{{ cookiecutter.app_name}}',
        '{{ cookiecutter.repo_name}}'
    ]

    # validate:
    for name in names:
        if not re.match(regex, name):
            return false
    return true


if __name__ == '__main__':
    # validate:
    if not validate_name(MODULE_REGEX):
        print('ERROR: The project name (%s) is not a valid module name. Please do not use a - and use _ instead' % module_name)

        # Exit to cancel project
        sys.exit(1)
