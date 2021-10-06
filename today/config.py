from configparser import ConfigParser
import os
import pathlib
from collections import namedtuple

# Resolve path to included defaults
DEFAULTS = pathlib.Path(os.path.realpath(__file__)).parent.joinpath('defaults.conf')

TodayConfig = namedtuple('today', ['datefmt', 'notes', 'scratch'])
NotesSection = namedtuple('notes', ['ext', 'directory', 'editor', 'viewer', 'search'])
ScratchSection = namedtuple('scratch', ['directory', 'viewer', 'search'])

def load_config(path):
    parser = ConfigParser(interpolation=None)
    # Load editor from EDITOR env var
    parser.read_dict({
        'notes': {
            'editor': os.environ.get('EDITOR', ''),
            'viewer': os.environ.get('EDITOR', '')
        }
    })

    # Load config from files
    parser.read([DEFAULTS, path])

    return TodayConfig(
        datefmt=parser.get('today', 'datefmt'),
        notes=NotesSection(
            ext=parser.get('notes', 'ext'),
            directory=parser.get('notes', 'directory'),
            editor=parser.get('notes', 'editor'),
            viewer=parser.get('notes', 'viewer'),
            search=parser.get('notes', 'search')
        ),
        scratch=ScratchSection(
            directory=parser.get('scratch', 'directory'),
            viewer=parser.get('scratch', 'viewer'),
            search=parser.get('scratch', 'search')
        )
    )
