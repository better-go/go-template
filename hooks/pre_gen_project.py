import re
import sys

MODULE_REGEX = r'^[_a-zA-Z][_a-zA-Z0-9]+$'

module_name = '{{ cookiecutter.project_slug}}'


def validate_name(regex, name):
    if not re.match(regex, name):
        print(
            'ERROR: The project slug (%s) is not a valid Python module name. Please do not use a - and use _ instead' % module_name)
        return false
    return true


if __name__ == '__main__':
    # validate:
    if not validate_name(MODULE_REGEX, module_name):
        # Exit to cancel project
        sys.exit(1)
