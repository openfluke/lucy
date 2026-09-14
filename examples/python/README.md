# Python / Jupyter examples (`openfluke-lucy`)

```bash
# from repo root
PYTHONPATH=python/src python3 examples/python/01_build_lpd.py
PYTHONPATH=python/src python3 examples/python/02_full_export.py
PYTHONPATH=python/src python3 examples/python/03_serve_client.py

# Jupyter
cd examples/python
PYTHONPATH=../../python/src jupyter notebook lucy_tutorial.ipynb
# or: pip install pandas jupyter
```

Local binary: `python/src/lucy/bin/lucy` (from `./go/scripts/build-artifacts.sh`).
Published: `pip install openfluke-lucy`.
