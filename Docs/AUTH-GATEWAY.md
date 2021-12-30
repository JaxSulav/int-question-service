# Authentication for grpc-gateway api endpoints

* Grpc-gateway exposes these REST JSON apis, which are http based endpoints and can be intercepted through http handler function.

### Using handlers to intercept any incoming request for REST api
```
    gwServer := &http.Server{
		Addr: gwRestPort,
		// Handle authentication through auth interceptor
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			bearer := r.Header.Get("Authorization")
			// Call grpc auth server
			err := GatewayAuthenticate(bearer)
			if err == nil {
				mux.ServeHTTP(w, r)
				return
			}
			
			// Case: Invalid auth token, write message to response writer object
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			err = json.NewEncoder(w).Encode(map[string]string{"code": "16", "message": "Unauthorized User"})
			if err != nil {
				log.Printf("Error writing to response writer: %v", err)
				return
			}
		}),
	}
	
	log.Fatalln(gwServer.ListenAndServe())
```

* The request returns with 401 Unauthorized error in case of invalid or missing token. 