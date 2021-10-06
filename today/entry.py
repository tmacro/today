import argparse
import os
import pathlib
import sys

from .config import load_config
from .tn.entry import build_parser as build_notes_parser

def path_type(path):
    return pathlib.Path(os.path.expanduser(path)).resolve()

def get_args():
    parser = argparse.ArgumentParser(
        prog=os.path.basename(sys.argv[0]),
        description='Simple scratch directory and notes manager.',
        formatter_class=argparse.ArgumentDefaultsHelpFormatter)

    parser.add_argument('-c', '--config',
        default='~/.config/today/today.conf',
        help='Specify an alternate config file')
    util_subparsers = parser.add_subparsers(title='Subcommands')
    note_subparser = util_subparsers.add_parser('note', help='use note commands')

    build_notes_parser(note_subparser)

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
