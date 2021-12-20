<?php

namespace App\Repositories;

use App\Models\User;
use Illuminate\Support\Facades\Hash;
use App\Traits\UtilsTrait;
use Illuminate\Support\Facades\Session;

class UserRepository{

    use UtilsTrait;

    public function login($data) {

        $user = User::where('mail', $data['mail'])->first();

        if ($user) {
            if (Hash::check($data['pass'], $user['pass']) == 1) {

                $user['is_blocked'] = true;

                $user->save();

                return $user['uuid'];
            } else {
                return "password don't match";
            }
        } else {
            return "user not found" ;
        }

    }

    public function register($data)
    {
        $data = [
            'name'=>$data['name'],
            'uuid'=>self::generteUUID(),
            'surnames'=>$data['surnames'],
            'mail'=>$data['mail'],
            'pass'=>Hash::make($data['pass']),
            'img'=>"https://upload.wikimedia.org/wikipedia/commons/thumb/7/7c/User_font_awesome.svg/2048px-User_font_awesome.svg.png",
            'google2fa_secret'=> null,
            'role' => 'client',
            'is_blocked' => false
        ];
        $user = User::create($data);
        return $user;
    }

    public function getUser() {
        $token = self::decode(session('token'));

        $array = json_decode(json_encode($token), True);

        $uuid = $array['uuid'];

        $user = User::where('uuid', $uuid)->first();

        return $user;
    }
}