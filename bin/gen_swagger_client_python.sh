# configuration
input="api-server.json"
output="tests/swagger_python_client"
type="python"
pkg_name="test_cli"
pkg_version="1.1.0"

source bin/gen_swagger_lib.sh
generate_swagger_lib $input $output $type $pkg_name $pkg_version

# update client package desc
sed -e '$ d' $output/setup.py >$output/setup.py.tmp
sed -e '$ d' $output/setup.py.tmp >$output/setup.py
sed -e '$ d' $output/setup.py >$output/setup.py.tmp
sed -e '$ d' $output/setup.py.tmp >$output/setup.py
echo """
    long_description_content_type='text/markdown',
    long_description=open('README.md', encoding='utf-8').read(),
)""" >>$output/setup.py
echo "

'''
Updating python api_client to use threadpoolexecutor with 10 max workers, removes runaway thread creation
issue with multiple async requests
'''  # noqa: E501
from concurrent.futures import ThreadPoolExecutor

def replace_pool_decorator(__init__):
    def new__init__(self, *args, **kw):
        __init__(self, *args, **kw)
        self._pool = ThreadPoolExecutor(max_workers=10)
    return new__init__

def mock_thread_pool_api(submit):
    def new_fn_apply_async(self, call_api, args):
        return self.submit(call_api, *args)
    return new_fn_apply_async

def mock_close(self):
    if self._pool:
        # 'ThreadPoolExecutor' object has no attribute 'close'
        if hasattr(self._pool, 'close'):
            self._pool.close()
            self._pool.join()
        self._pool = None
        if hasattr(atexit, 'unregister'):
            atexit.unregister(self.close)

ApiClient.__init__ = replace_pool_decorator(ApiClient.__init__)
ApiClient.close = mock_close

ThreadPoolExecutor.apply_async = mock_thread_pool_api(ThreadPoolExecutor.submit)

" >> $output/$pkg_name/api_client.py

echo "
    validation_enabled: bool = True  # default is True
" >> $output/$pkg_name/configuration.py


  if [ -f $output/$pkg_name/__main__.py ]; then
    sed -e 's/pythonic_params=True/pythonic_params=False/' $output/$pkg_name/__main__.py > $output/$pkg_name/__main__.py.tmp
    mv -f $output/$pkg_name/__main__.py.tmp $output/$pkg_name/__main__.py
  fi
  # make sed work on both macos and linux https://stackoverflow.com/a/14946954
  sed -e 's/install_requires=REQUIRES/install_requires=REQUIRES + ["querystring==0.1.0"]/g' $output/setup.py > $output/setup.py.tmp
  mv -f $output/setup.py.tmp $output/setup.py

  # Move "swagger_server/swagger_client" to the root directory so that we could import them easily.
  # And we don't need extra step to tell IDE how to locate source code directory.
#   rm -rf $pkg_name
#   cp -rv $output/$pkg_name $pkg_name
  #rm -rf $output
  echo  '## Demo 
``` python
import test_cli
from test_cli.models import Cypher, CypherOperation

configuration = test_cli.Configuration()
configuration.host = "https://test.cn"
# Enter a context with an instance of the API client
with test_cli.ApiClient(configuration) as api_client:
    api_client.set_default_header(
        "Authorization", "APPCODE test")

    api_instance = test_cli.PlatformApi(api_client)
    cypher = Cypher(store_id="platform.test",
                    operation=CypherOperation.DECRYPTBATCH,
                    items=["$Kg7PP/qJTnIYX+22Y2pa+A==$1$"],
                    platform="platform",
                    )
    try:
        api_response = api_instance.batch_operate_cypher(cypher=cypher)
        print(api_response, api_response.data)
    except Exception as e:
        print("Exception : %s\n" % str(e))
```
' >> $output/README.md

  echo "${green}Done!${reset}"