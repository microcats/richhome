package config

const (
    LogFormat = `{"time":"${time_rfc3339_nano}","id":"${id}","remote_ip":"${remote_ip}","host":"${host}",` +
        `"method":"${method}","uri":"${uri}","status":${status}, "latency":${latency},` +
        `"latency_human":"${latency_human}","bytes_in":${bytes_in},` +
        `"bytes_out":${bytes_out}}` + "\n"
)
