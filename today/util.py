import datetime
import pathlib
import os.path

from .constant import TODAY

def get_datestamp(fmt, date=None):
    if date is None:
        date = datetime.date.today()
    return date.strftime(fmt)

def resolve_path(path):
    return pathlib.Path(os.path.expanduser(path)).resolve()

def path_type(path):
    return pathlib.Path(os.path.expanduser(path)).resolve()

def relative_date_type(arg):
    days = int(arg)
    return TODAY + datetime.timedelta(days=days)

def date_type(arg):
    try:
        date = datetime.datetime.strptime(arg, '%m.%d').date()
        return date.replace(year=TODAY.year)
    except ValueError:
        pass
    try:
        date = datetime.datetime.strptime(arg, '%d').date()
        return date.replace(year=TODAY.year, month=TODAY.month)
    except ValueError:
        pass
    return datetime.datetime.strptime(arg, '%Y.%m.%d').date()

def date_or_relative(arg):
    try:
        return date_type(arg)
    except ValueError:
        pass
    return relative_date_type(arg)

class CommandRegistry:
    def __init__(self):
        self._commands = dict()

    def get(self, name):
        return self._commands[name]

    def register(self, name):
        def inner(func):
            if name in self._commands:
                raise ValueError(f'Command already registered: {name}')
            self._commands[name] = func
            return func
        return inner


