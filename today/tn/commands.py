import subprocess
import os
from datetime import datetime

from ..util import get_datestamp, resolve_path, CommandRegistry

commands = CommandRegistry()

def get_notes_dir(config):
    return resolve_path(config.notes.directory)

def get_filename(config, date):
    return resolve_path(f'{get_notes_dir(config)}/{get_datestamp(config.datefmt, date=date)}{config.notes.ext}')

@commands.register('create')
def create_note(args, config):
    filename = get_filename(config, None)
    if not filename.parent.exists():
        filename.parent.mkdir(parents=True)
    if not filename.exists():
        with open(filename.as_posix(), 'w') as f:
            f.write(f'# {get_datestamp(config.datefmt)}\n')

@commands.register('edit')
def edit_note(args, config):
    filename = get_filename(config, args.day)
    if not filename.exists():
        create_note(args, config)
    subprocess.run(f'{config.notes.editor} {filename}', shell=True, env=os.environ)


@commands.register('view')
def view_note(args, config):
    filename = get_filename(config, args.day)
    subprocess.run(f'{config.notes.viewer} {filename}', shell=True, env=os.environ)

@commands.register('search')
def search_notes(args, config):
    directory = get_notes_dir(config)
    subprocess.run(f'{config.notes.search} {args.expression} {directory}', shell=True, env=os.environ)

@commands.register('list')
def list_notes(args, config):
    directory = get_notes_dir(config)
    for note in sorted(directory.glob(f'*{config.notes.ext}')):
        print(note)

@commands.register('prune')
def prune_notes(args, config):
    directory = get_notes_dir(config)
    for note in directory.glob(f'*{config.notes.ext}'):
        with open(note) as f:
            lines = f.readlines()
        if len(lines) == 0:
            print(note)
            if not args.dry_run:
                note.unlink()
        elif len(lines) == 1:
            parts = lines[0].strip().split(' ')
            if len(parts) == 2:
                try:
                    datetime.strptime(parts[1], config.datefmt)
                    print(note)
                    if not args.dry_run:
                        note.unlink()
                except ValueError as e:
                    pass
