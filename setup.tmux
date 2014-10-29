#!/bin/bash -x

# This shell scripts sets up the tmux environment for me. 

SESSION_NAME="hipchat_tools"

setup_new_session() {
	 #function_body
	 tmux new-session -s $SESSION_NAME -n editor -d
	 tmux send-keys -t $SESSION_NAME 'vim' C-m
	 tmux send-keys -t $SESSION_NAME ':NERDTree' C-m
	 tmux split-window -v -p 10 -t $SESSION_NAME
	 tmux send-keys -t $SESSION_NAME:1.2 'cd ../../../..' C-m
	 tmux send-keys -t $SESSION_NAME:1.2 'clear' C-m
	 tmux new-window -n console -t $SESSION_NAME
	 tmux send-keys -t $SESSION_NAME:2 'cd ../../../..' C-m
	 tmux send-keys -t $SESSION_NAME:2 'clear' C-m
}

tmux list-sessions | grep $SESSION_NAME
E=$?
if [ $E != 0 ];then
setup_new_session
fi
tmux select-window -t $SESSION_NAME:1
tmux select-pane -t $SESSION_NAME:1.0
tmux attach -t $SESSION_NAME


