import setuptools

setuptools.setup(
    name='today',
    version='0.1.0',
    entry_points={
        'console_scripts': [
            'tn=today.tn.entry:cli',
            'td=today.td.entry:cli',
            'today=today.entry:cli',
        ]
    }
)
