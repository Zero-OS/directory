# Directory Test Suite

## Exectuion steps:
```bash
git clone https://github.com/zero-os/0-directory.git
cd 0-directory/clients/pyclient
pip3 install .
cd ../../testsuite/
pip3 install -r requirements.txt
export PYTHONPATH='./'
nosetests-3.4 -s -v --logging-level=WARNING --progressive-with-bar --rednose test_cases --tc-file=config.ini
```

note: You have to update the config.ini file

