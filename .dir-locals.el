((go-mode . ((dape-configs . ((go-debug-redditview
                               modes (go-mode go-ts-mode)
                               command "dlv"
                               command-args ("dap" "--listen" "127.0.0.1::autoport")
                               command-cwd dape-command-cwd
                               host "127.0.0.1"
                               port :autoport
                               :request "launch"
                               :mode "debug"
                               :type "go"
                               :showLog "true"
                               :cwd "."
                               :args ["https://www.reddit.com/r/electronics/comments/1k6zst4/showcase_my_finished_digital_oscilloscope_project/"]
                               :program "./cmd/redditview"))))))
