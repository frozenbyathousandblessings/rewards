function validateRewards(){
    var mobile = document.getElementById("mobile");
       axios.post('https://fgep4oij3l.execute-api.ap-southeast-2.amazonaws.com/Prod/verifyrewards', {
       headers: {
       	  'Access-Control-Allow-Origin': '*',
       	},
        Mobilenumber: mobile.value
      })
      .then(function (res) {
        if(res.data.Status==200){
            document.getElementById('rewardsres').innerHTML='Rewards Redeemed For '+res.data.MobileNumber;
            document.getElementById('rewardsres').setAttribute('style','text-align:center;display:block;color:white;font-weight:bold');

        }
        else{
            document.getElementById('rewardsres').innerHTML='Invalid data for Rewards';
            document.getElementById('rewardsres').setAttribute('style','text-align:center;display:block;color:red;font-weight:bold');
        }

      })
      .catch(function (error) {
                alert('Error when checking rewards'+error)
      });
    }