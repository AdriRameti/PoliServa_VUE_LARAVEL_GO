<?php

namespace App\Http\Requests;

use Illuminate\Foundation\Http\FormRequest;
use Illuminate\Validation\ValidationException;
use Illuminate\Contracts\Validation\Validator;
use Illuminate\Http\Exceptions\HttpResponseException;
use App\Models\User;

class RegisterRequest extends FormRequest
{
    /**
     * Determine if the user is authorized to make this request.
     *
     * @return bool
     */
    public function authorize()
    {
        return true;
    }

    /**
     * Get the validation rules that apply to the request.
     *
     * @return array
     */
    public function rules()
    {
        return [
            'name' => 'required|string',
            'surnames' => 'required|string',
            'mail' => 'required|max:255|email|unique:users',
            'pass' =>'required',
            'social'=>'required'
        ];
    }
    public function messages(){
        return [
            'name.required'     => 'User name is missing!',
            'surnames.required'     => 'User name is missing!',
            'mail.required'    => 'Email address is missing!',
            'mail.unique'      => 'This email has been used already!',
            'pass.required' => 'Password is missing',
            'social.required' => 'Social is missing',
        ];
    }
}
