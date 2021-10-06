import setuptools

setuptools.setup(
    name='today',
    version='0.1.0',
    description='Simple scratch directory and notes manager',
    author='Taylor McKinnon',
    author_email='mail@tmacs.space',
    url='https://github.com/tmacro/today/',
    packages=['today', 'today.tn', 'today.td'],
    package_data={
        'today': [ 'defaults.conf']
    },
    entry_points={
        'console_scripts': [
            'tn=today.tn.entry:cli',
            'td=today.td.entry:cli',
            'today=today.entry:cli',
        ]
    }
)
