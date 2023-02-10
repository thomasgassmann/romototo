import logging
import click
import yaml
import os
import os.path
from romototo.runner import run_with_config


def get_log_level(level_str: str) -> int:
    try:
        import pydevd # if debugger is attached, set log level to debug
        return logging.DEBUG
    except ImportError:
        try:
            return getattr(logging, level_str) if level_str else logging.INFO
        except AttributeError:
            raise click.ClickException(f'Log level {level_str} not found.')

def setup_logger(level_str: str):
    log_formatter = logging.Formatter(
        "%(asctime)s [%(threadName)-12.12s] [%(levelname)-5.5s]  %(message)s")
    root_logger = logging.getLogger()
    root_logger.propagate = True
    root_logger.setLevel(get_log_level(level_str))

    console_handler = logging.StreamHandler()
    console_handler.setFormatter(log_formatter)
    root_logger.addHandler(console_handler)

@click.group()
@click.option('--log-level', nargs=1, required=False)
def romototo_cli(log_level: str):
    setup_logger(log_level)

@click.command(name='exec')
@click.argument('config')
def exec_cmd(config: str):
    if not os.path.isfile(config):
        raise click.ClickException(f'Config file {config} not found.')
    with open(config, 'r') as stream:
        try:
            run_with_config(yaml.safe_load(stream))
        except yaml.YAMLError as exc:
            raise click.ClickException(f'Invalid file {config}')

romototo_cli.add_command(exec_cmd)

def main():
    romototo_cli(obj={})

if __name__ == '__main__':
    main()
