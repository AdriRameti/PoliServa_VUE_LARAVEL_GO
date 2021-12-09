<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;

class Sport extends Model
{
    use HasFactory;

    protected $fillable = ['slug', 'name', 'img'];
    protected $hidden = ['created_at', 'updated_at'];


    public function courts() {
        return $this->hasMany(Court::class);
    }
}
