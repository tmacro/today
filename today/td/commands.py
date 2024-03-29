import subprocess
import os
from datetime import datetime

from ..util import get_datestamp, resolve_path, CommandRegistry
from ..constant import TODAY

commands = CommandRegistry()

def get_scratch_dir(config):
    return resolve_path(config.scratch.directory)

def get_dir_name(config, date):
    return resolve_path(f'{get_scratch_dir(config)}/{get_datestamp(config.datefmt, date=date)}')

@commands.register('create')
def create_dir(args, config):
    dir_name = get_dir_name(config, None)
    if not dir_name.exists():
        dir_name.mkdir(parents=True)

@commands.register('view')
def view_dir(args, config):
    dir_name = get_dir_name(config, args.day)
    if not dir_name.exists():
        dir_name.mkdir()
    cmd = config.scratch.viewer.format(directory=dir_name)
    subprocess.run(cmd, shell=True, env=os.environ)

@commands.register('search')
def search_dirs(args, config):
    directory = get_scratch_dir(config)
    cmd = config.scratch.search.format(directory=directory, expression=args.expression)
    subprocess.run(cmd, shell=True, env=os.environ)

def _is_scratch_dir(config, path):
    try:
        datetime.strptime(path.name, config.datefmt)
    except ValueError as e:
        return False
    return True

@commands.register('list')
def list_dirs(args, config):
    directory = get_scratch_dir(config)
    for path in sorted(directory.iterdir()):
        if not path.is_dir() or not _is_scratch_dir(config, path):
            continue
        print(path)

@commands.register('prune')
def prune_dirs(args, config):
    directory = get_scratch_dir(config)
    for path in sorted(directory.iterdir()):
        if not path.is_dir() or not _is_scratch_dir(config, path):
            continue
        contents = list(path.iterdir())
        if len(contents) > 1:
            continue
        if len(contents) == 1 and contents[0].name != 'NOTES.md':
            continue
        print(path)
        if not args.dry_run:
            path.rmdir()

@commands.register('show')
def show_dir(args, config):
    dir_name = get_dir_name(config, args.day)
    if args.day == TODAY and not dir_name.exists():
        dir_name.mkdir(parents=True)
    print(dir_name)

@commands.register('find')
def find_in_dirs(args, config):
    directory = get_scratch_dir(config)
    cmd = config.scratch.find.format(directory=directory, expression=args.expression)
    subprocess.run(cmd, shell=True, env=os.environ)
