<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;
use Illuminate\Foundation\Auth\User as Authenticatable;
use Illuminate\Support\Facades\Hash;

class User extends Authenticatable
{
    use HasFactory;
    protected $fillable = ['name', 'surnames','uuid', 'mail', 'pass', 'img', 'google2fa_secret', 'role', 'is_blocked'];
    protected $hidden = ['created_at', 'updated_at', 'google2fa_secret'];

    public function getJWTIdentifier()
    {
        return $this->getKey();
    }

    public function getJWTCustomClaims()
    {
        return [];
    }

    /**
     * Ecrypt the user's google_2fa secret.
     *
     * @param  string  $value
     * @return string
     */

    // public function setGoogle2faSecretAttribute($value)
    // {
    //     $this->attributes['google2fa_secret'] = encrypt($value);
    // }

    /**
     * Decrypt the user's google_2fa secret.
     *
     * @param  string  $value
     * @return string
     */
    public function getGoogle2faSecretAttribute($value)
    {
        return decrypt($value);
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
