package com.fzflabs.aiboilerplate.android.repository

import com.fzflabs.aiboilerplate.android.model.ProductDetailsDto
import com.fzflabs.aiboilerplate.android.network.ApiService
import com.fzflabs.aiboilerplate.android.util.ApiState
import javax.inject.Inject

class RepositoryImpl @Inject constructor(
    private val apiService: ApiService
) : Repository {

    override suspend fun getProductDetails(): ApiState<ProductDetailsDto> = try {
        ApiState.Success(apiService.getProductDetails())
    } catch (e: Exception) {
        ApiState.Error(errorMsg = e.message.toString())
    }
}