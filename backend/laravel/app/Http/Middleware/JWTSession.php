<?php

namespace App\Http\Middleware;

use Illuminate\Support\Facades\Session;
use Closure;

class JWTSession
{
    
    public function handle($request, Closure $next) {

        $token = $request->bearerToken();

        if ($token) {
            session(['token'=> $token]);
            Session::save();
        }

        return $next($request);
    }
}