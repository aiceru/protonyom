
?
protonyom_models.proto	protonyom"
EmptyParams"1
	OAuthInfo
id (	Rid
email (	Remail"?
Account
id (	Rid
name (	Rname
email (	Remail 
hasPassword (RhasPassword?
	oauthinfo (2!.protonyom.Account.OauthinfoEntryR	oauthinfo
photourl (	Rphotourl
signedup (Rsignedup
pets (	RpetsR
OauthinfoEntry
key (	Rkey*
value (2.protonyom.OAuthInfoRvalue:8"?
Pet
id (	Rid
name (	Rname
photourl (	Rphotourl
adopted (Radopted
family (	Rfamily
species (	Rspecies
feeders (	Rfeeders"?
Feed
id (	Rid
petId (	RpetId
	timestamp (R	timestamp
feederId (	RfeederId

feederName (	R
feederName
amount (Ramount
unit (	RunitB$Z"github.com/aiceru/protonyom/gonyomJ?
  *

  

 

 9
	
 9
	
  


 


 





 

 

 	

 

	

	

		

	
7
 + login 시 app 에 보내줘야 할 정보





 

 

 	

 





	







	











'



"

%&





	


&
" timestamp in epoch sec.




















  




 

 

 	

 





	







	


&
" timestamp in epoch sec.












	







	















" *


"

 #

 #

 #	

 #

$

$

$	

$
&
%" timestamp in epoch sec.


%

%

%

&

&

&	

&

'

'

'	

'

(

(

(	

(

)

)

)	

)bproto3
?
protonyom_api_account.proto	protonyomprotonyom_models.proto"
GetAccountRequest"?
GetAccountReply,
account (2.protonyom.AccountRaccount"@
UpdateAccountRequest
path (	Rpath
value (	Rvalue"B
UpdateAccountReply,
account (2.protonyom.AccountRaccount"&
DeleteAccountRequest
id (	Rid"
DeleteAccountReply"+
AcceptInviteRequest
petId (	RpetId"A
AcceptInviteReply,
account (2.protonyom.AccountRaccount"j
UploadProfileRequest"
profilePhoto (RprofilePhoto.
profileContentType (	RprofileContentType"E
UploadProfileResponse,
account (2.protonyom.AccountRaccount2?

AccountApiA
Get.protonyom.GetAccountRequest.protonyom.GetAccountReply" J
Update.protonyom.UpdateAccountRequest.protonyom.UpdateAccountReply" J
Delete.protonyom.DeleteAccountRequest.protonyom.DeleteAccountReply" N
AcceptInvite.protonyom.AcceptInviteRequest.protonyom.AcceptInviteReply" T
UploadProfile.protonyom.UploadProfileRequest .protonyom.UploadProfileResponse" B$Z"github.com/aiceru/protonyom/gonyomJ?
  5

  

 

 9
	
 9
	
   


  


 

  9

  	

  


  &5

 	B

 	

 	!

 	,>

 
B

 


 
!

 
,>

 F

 

 &

 1B

 L

 

 (

 3H


  


 


 




 

 	

 


 


 




 

 

 	

 





	




 




 

 	

 


 


 !




  

  

  	

  


# $


#


& (


&

 '

 '

 '	

 '


* ,


*

 +

 +	

 +


 +


. 1


.

 /

 /

 /

 /

0 

0

0	

0


	3 5


	3

	 4

	 4	

	 4


	 4bproto3
?
protonyom_api_feed.proto	protonyomprotonyom_models.proto"5
AddFeedRequest#
feed (2.protonyom.FeedRfeed"3
AddFeedReply#
feed (2.protonyom.FeedRfeed"]
GetFeedsRequest
petId (	RpetId

startAfter (R
startAfter
limit (Rlimit"6
GetFeedsReply%
feeds (2.protonyom.FeedRfeeds"A
DeleteFeedRequest
petId (	RpetId
feedId (	RfeedId"
DeleteFeedReply"8
UpdateFeedRequest#
feed (2.protonyom.FeedRfeed"6
UpdateFeedReply#
feed (2.protonyom.FeedRfeed2?
FeedApi?
AddFeed.protonyom.AddFeedRequest.protonyom.AddFeedReply" B
GetFeeds.protonyom.GetFeedsRequest.protonyom.GetFeedsReply" H

DeleteFeed.protonyom.DeleteFeedRequest.protonyom.DeleteFeedReply" H

UpdateFeed.protonyom.UpdateFeedRequest.protonyom.UpdateFeedReply" B$Z"github.com/aiceru/protonyom/gonyomJ?
  *

  

 

 9
	
 9
	
   


  


 

  7

  

  

  '3

 	:

 	

 	

 	)6

 
@

 


 
"

 
-<

 @

 

 "

 -<


  


 

  

  

  

  


 




 

 

 

 


 




 

 

 	

 
"
" last feed timestamp

















 




 

 


 

 

 


 !




 

 

 	

 

 

 

 	

 


" #


"


% '


%

 &

 &

 &

 &


( *


(

 )

 )

 )

 )bproto3
?
protonyom_api_pet.proto	protonyomprotonyom_models.proto"?
Family
name (	Rname8
species (2.protonyom.Family.SpeciesEntryRspecies:
SpeciesEntry
key (	Rkey
value (	Rvalue:8"0
GetFamiliesRequest
language (	Rlanguage"?
GetFamiliesReplyE
families (2).protonyom.GetFamiliesReply.FamiliesEntryRfamiliesN
FamiliesEntry
key (	Rkey'
value (2.protonyom.FamilyRvalue:8"?
AddPetRequest 
pet (2.protonyom.PetRpet"
profilePhoto (RprofilePhoto.
profileContentType (	RprofileContentType"_
AddPetReply,
account (2.protonyom.AccountRaccount"
pets (2.protonyom.PetRpets"?
UpdatePetRequest 
pet (2.protonyom.PetRpet"
profilePhoto (RprofilePhoto.
profileContentType (	RprofileContentType"4
UpdatePetReply"
pets (2.protonyom.PetRpets"(
DeletePetRequest
petId (	RpetId"b
DeletePetReply,
account (2.protonyom.AccountRaccount"
pets (2.protonyom.PetRpets"+
GetPetListRequest
petIds (	RpetIds"5
GetPetListReply"
pets (2.protonyom.PetRpets"%
GetPetRequest
petId (	RpetId"/
GetPetReply 
pet (2.protonyom.PetRpet2?
PetApiK
GetFamilies.protonyom.GetFamiliesRequest.protonyom.GetFamiliesReply" <
AddPet.protonyom.AddPetRequest.protonyom.AddPetReply" E
	UpdatePet.protonyom.UpdatePetRequest.protonyom.UpdatePetReply" E
	DeletePet.protonyom.DeletePetRequest.protonyom.DeletePetReply" H

GetPetList.protonyom.GetPetListRequest.protonyom.GetPetListReply" <
GetPet.protonyom.GetPetRequest.protonyom.GetPetReply" B$Z"github.com/aiceru/protonyom/gonyomJ?
  C

  

 

 9
	
 9
	
   


  


 

  C

  

  $

  /?

 	4

 	

 	

 	%0

 
=

 


 
 

 
+9

 =

 

  

 +9

 @

 

 "

 -<

 4

 

 

 %0


  


 

  

  

  	

  

 "

 

 

  !


 




 

 

 	

 


 




 #

 

 

 !"


  




 

 

 	

 









 



	




! $


!

 "

 "	

 "


 "

#

#


#

#

#


& *


&

 '

 '

 '	

 '

(

(

(

(

) 

)

)	

)


+ -


+

 ,

 ,


 ,

 ,

 ,


/ 1


/

 0

 0

 0	

 0


2 5


2

 3

 3	

 3


 3

4

4


4

4

4


	7 9


	7

	 8

	 8


	 8

	 8

	 8



: <



:


 ;


 ;



 ;


 ;


 ;


> @


>

 ?

 ?

 ?	

 ?


A C


A

 B

 B

 B	

 Bbproto3
?
protonyom_api_sign.proto	protonyomprotonyom_models.proto"?
SignUpRequest
name (	Rname
email (	Remail
password (	H Rpassword4
	oauthinfo (2.protonyom.OAuthInfoH R	oauthinfo$
oauthprovider (	Roauthprovider
photourl (	RphotourlB

credential"=
	EmailCred
email (	Remail
password (	Rpassword"?
SignInRequest4
	emailcred (2.protonyom.EmailCredH R	emailcred4
	oauthinfo (2.protonyom.OAuthInfoH R	oauthinfo$
oauthprovider (	RoauthproviderB

credential"O
	SignReply,
account (2.protonyom.AccountRaccount
token (	Rtoken2?
SignApi:
SignIn.protonyom.SignInRequest.protonyom.SignReply" :
SignUp.protonyom.SignUpRequest.protonyom.SignReply" ;
SignOut.protonyom.EmptyParams.protonyom.EmptyParams" B$Z"github.com/aiceru/protonyom/gonyomJ?
  (

  

 

 9
	
 9
	
   


  


 

  2

  

  

  %.

 	2

 	

 	

 	%.

 
3

 


 


 
$/


  


 

  

  

  	

  

 

 

 	

 

  

  

 

 


 

 

 

 

 

 

 

 

 	

 

 

 

 	

 


 




 

 

 	

 





	




 #




 !

 

 

 

 

 

 

 

 

 

"

"

"	

"


% (


%

 &

 &	

 &


 &

'

'

'	

'bproto3