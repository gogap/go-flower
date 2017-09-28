CONFIG, _ = CTX.Get("CONFIG")
CODE, _   = CTX.Get("CODE")
LOG, _    = CTX.Get("LOG")
REDIS, _  = CTX.Get("REDIS")
ARGS, _   = CTX.Get("ARGS")

CONFIG = CONFIG.GetConfig("devops.creation")



CanContinue = fn(flag) {
	if ARGS.Exist(flag) || ARGS.Exist("all") {
		return true
	}
	return false
}