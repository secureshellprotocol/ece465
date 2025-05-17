# Endpoint provision

Endpoint should be pre-provisioned with the `provision/` Ansible setup.

It should NOT have the env-var `STARTING` set at all!

This is just a quick configuration for a round-robin reverse proxy between our
worker nodes. For simplicity, we are statically defined.

See `caddy/conf/Caddyfile`.

## Details

Within `Vagrantfile`, we have an endpoint VM which sits between the User and our
internal workers{1..3}. 

To test: 
```
curl 192.168.56.100
```

Annotated `Caddyfile`:
```
*:80 {       # on inbound port 80
reverse_proxy <worker1_ip>:8000 <worker2_ip>:8000 <worker3_ip>:8000 {
        lb_policy round_robin       # go around each server, and take the first
                                    #   which responds.
        lb_policy retries 4         # try to round robin all servers, or fail if
                                    #   all are unavailable
    }
}
```

For debugging purposes, each worker can also be directly `curl`ed, to check for
functionality
