
:: 新窗口启动 user rpc
start cmd /k "d: && cd D:\GoProjects\kun_chat\user\rpc && go run ."
:: 新窗口启动 user api
start cmd /k "d: && cd D:\GoProjects\kun_chat\user\api && go run ."
:: 新窗口延迟 1 秒启动 relation rpc, 
start cmd /k "timeout -nobreak 1 && d: && cd D:\GoProjects\kun_chat\relation\rpc && go run ."

:: 新窗口延迟 1 秒启动 relation api
start cmd /k "timeout -nobreak 1 && d: && cd D:\GoProjects\kun_chat\relation\api && go run ."

:: 新窗口启动 chat api
start cmd /k "d: && cd D:\GoProjects\kun_chat\chat\api && go run ."
