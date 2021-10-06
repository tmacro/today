import datetime
import pathlib
import os.path

def get_datestamp(fmt, date=None):
    if date is None:
        date = datetime.date.today()
    return date.strftime(fmt)

def resolve_path(path):
    return pathlib.Path(os.path.expanduser(path)).resolve()


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


