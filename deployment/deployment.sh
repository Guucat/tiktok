alias k=kubectl

OP=apply

k $OP -f ./comment/comment.yaml
k $OP -f ./snowflake/snowflake.yaml
k $OP -f ./user/user.yaml
k $OP -f ./route/route.yaml