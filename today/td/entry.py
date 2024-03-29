import argparse
import configparser
import datetime
import os
import pathlib
import sys

from ..config import load_config
from ..util import path_type, date_or_relative
from .commands import commands

from ..constant import TODAY

def build_parser(parser):
    dir_subparsers = parser.add_subparsers(title='Scratch directory commands')
    create_dir_subparser = dir_subparsers.add_parser('create', help='Create an empty scratch directory.')
    create_dir_subparser.set_defaults(command=commands.get('create'))

    show_dir_subparser = dir_subparsers.add_parser('show', help='Print the path of the scratch directory.')
    show_dir_subparser.set_defaults(command=commands.get('show'))
    show_dir_subparser.add_argument('day', nargs='?', default=TODAY, type=date_or_relative, help='Specify a date relative to today')

    view_dir_subparser = dir_subparsers.add_parser('view', help="View a scratch directory's contents.")
    view_dir_subparser.set_defaults(command=commands.get('view'))
    view_dir_subparser.add_argument('day', nargs='?', default=TODAY, type=date_or_relative, help='Specify a date relative to today')

    list_dir_subparser = dir_subparsers.add_parser('list', help='List scratch directories.')
    list_dir_subparser.set_defaults(command=commands.get('list'))

    search_dir_subparser = dir_subparsers.add_parser('search', help='Search scratch directories.')
    search_dir_subparser.set_defaults(command=commands.get('search'))
    search_dir_subparser.add_argument('expression', help='Search dirs using a regex')

    prune_dir_subparser = dir_subparsers.add_parser('prune', help='Remove empty scratch directories.')
    prune_dir_subparser.set_defaults(command=commands.get('prune'))
    prune_dir_subparser.add_argument('-d', '--dry-run', action='store_true', help='Only print filenames detected for deletion without deleting')

    find_dir_subparser = dir_subparsers.add_parser('find', help='Search scratch directories by filename')
    find_dir_subparser.set_defaults(command=commands.get('find'))
    find_dir_subparser.add_argument('expression', help='Search dirs for a file')

def get_args():
    parser = argparse.ArgumentParser(
        prog=os.path.basename(sys.argv[0]),
        description='Simple scratch directory manager',
        formatter_class=argparse.ArgumentDefaultsHelpFormatter)

    parser.add_argument('-c', '--config',
        default='~/.config/today/today.conf',
        type=path_type,
        help='Specify an alternate config file')

    build_parser(parser)

    args = parser.parse_args()
    if not hasattr(args, 'command'):
        return parser.parse_args(['--help'])
    return args

def cli():
    args = get_args()
    config = load_config(args.config)
    args.command(args, config)
