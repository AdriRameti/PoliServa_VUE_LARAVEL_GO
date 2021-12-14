<?php

namespace App\Repositories;

use App\Models\User;
use Illuminate\Support\Facades\Hash;
use App\Traits\UtilsTrait;

class UserRepository{

    use UtilsTrait;

    public function register($data)
    {
        $data = [
            'name'=>$data['name'],
            'uuid'=>self::generteUUID(),
            'surnames'=>$data['surnames'],
            'mail'=>$data['mail'],
            'pass'=>Hash::make($data['pass']),
            'img'=>"https://upload.wikimedia.org/wikipedia/commons/thumb/7/7c/User_font_awesome.svg/2048px-User_font_awesome.svg.png"
        ];
        $user = User::create($data);
        return $user;
    }
}