<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;
use Illuminate\Foundation\Auth\User as Authenticatable;
use Illuminate\Support\Facades\Hash;

class User extends Authenticatable
{
    use HasFactory;
    protected $fillable = ['name', 'surnames','uuid', 'mail', 'pass', 'img'];
    protected $hidden = ['created_at', 'updated_at'];

    public function getJWTIdentifier()
    {
        return $this->getKey();
    }

    public function getJWTCustomClaims()
    {
        return [];
    }
    public function setPassAttribute($value)
    {
        $this->attributes['pass'] = Hash::make($value);
    }
    public function getFullNameAttribute()
    {
        return $this->attributes['name'] .' '. $this->attributes['surnames'];
    }
}
