let s:channel = ch_open("localhost:80", {"mode": "raw"})

func! MyHandler(channel, msg) abort
    new RESPONSE | set buftype=nofile
    call setline(".", a:msg->split("\r\n"))
    nnoremap <buffer> <silent> q :bw!<CR>
endfunc

let s:get_request = "GET / HTTP/1.1\r\nHost: localhost\r\n\r\n"

call ch_sendraw(s:channel, s:get_request, {'callback': "MyHandler"})
