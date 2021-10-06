import argparse
import configparser
import datetime
import os
import pathlib
import sys

from ..config import load_config
from .commands import commands

TODAY = datetime.date.today()

def path_type(path):
    return pathlib.Path(os.path.expanduser(path)).resolve()

def relative_date_type(arg):
    days = int(arg)
    return TODAY + datetime.timedelta(days=days)

def build_parser(parser):
    note_subparsers = parser.add_subparsers(title='Note Commands')
    create_note_subparser = note_subparsers.add_parser('create', help='Create an empty note.')
    create_note_subparser.set_defaults(command=commands.get('create'))

    edit_note_subparser = note_subparsers.add_parser('edit', help='Edit a note.')
    edit_note_subparser.set_defaults(command=commands.get('edit'))
    edit_note_subparser.add_argument('day', nargs='?', default=TODAY, type=relative_date_type, help='Specify a date relative to today')

    view_note_subparser = note_subparsers.add_parser('view', help='View a note.')
    view_note_subparser.set_defaults(command=commands.get('view'))
    view_note_subparser.add_argument('day', nargs='?', default=TODAY, type=relative_date_type, help='Specify a date relative to today')

    list_note_subparser = note_subparsers.add_parser('list', help='List notes.')
    list_note_subparser.set_defaults(command=commands.get('list'))

    search_note_subparser = note_subparsers.add_parser('search', help='Search notes.')
    search_note_subparser.set_defaults(command=commands.get('search'))
    search_note_subparser.add_argument('expression', help='Search notes using a regex')

    prune_note_subparser = note_subparsers.add_parser('prune', help='Remove empty notes.')
    prune_note_subparser.set_defaults(command=commands.get('prune'))
    prune_note_subparser.add_argument('-d', '--dry-run', action='store_true', help='Only print filenames detected for deletion without deleting')

def get_args():
    parser = argparse.ArgumentParser(
        prog=os.path.basename(sys.argv[0]),
        description='Simple notes manager',
        formatter_class=argparse.ArgumentDefaultsHelpFormatter)

    parser.add_argument('-c', '--config',
        default='~/.config/today/today.conf',
        help='Specify an alternate config file')

    build_parser(parser)

    args = parser.parse_args()
    print(args)
    if not hasattr(args, 'command'):
        return parser.parse_args(['--help'])
    return args

def cli():
    args = get_args()
    config = load_config(args.config)
    print(config)
    args.command(args, config)
