CONFIG, _ = CTX.Get("CONFIG")
CODE, _   = CTX.Get("CODE")
LOG, _    = CTX.Get("LOG")
REDIS, _  = CTX.Get("REDIS")
ARGS, _   = CTX.Get("ARGS")

CONFIG = CONFIG.GetConfig("devops.creation")



ShouldExecute = fn(flag) {
	if ARGS.Exist(flag) || ARGS.Exist("all") {
		return true
	}
	return false
}


SetENV = fn(key, value) {
	err = REDIS.HSet(CODE+"::ENVS", key, value).Err()

	if err!=nil {
		LOG.WithField("CODE", CODE).
			WithField("KEY", key).
			Warnln("Env fail to set")
	} else {
		LOG.WithField("CODE", CODE).
			WithField("KEY", key).
			Debugln("Env setted")
	}


	return err
}

GetENV = fn(key) {

	LOG.WithField("KEY",key).Debug("Get env")

	if !REDIS.HExists(CODE+"::ENVS", key).Val() {
		return "", false
	}

	result = REDIS.HGet(CODE+"::ENVS", key).Val()
 
	return result, true
}